// Copyright 2020-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufcli

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/bufbuild/buf/private/buf/bufwkt/bufwktstore"
	"github.com/bufbuild/buf/private/bufpkg/bufmodule"
	"github.com/bufbuild/buf/private/bufpkg/bufmodule/bufmoduleapi"
	"github.com/bufbuild/buf/private/bufpkg/bufmodule/bufmodulecache"
	"github.com/bufbuild/buf/private/bufpkg/bufmodule/bufmodulestore"
	"github.com/bufbuild/buf/private/bufpkg/bufplugin"
	"github.com/bufbuild/buf/private/bufpkg/bufplugin/bufpluginapi"
	"github.com/bufbuild/buf/private/bufpkg/bufplugin/bufplugincache"
	"github.com/bufbuild/buf/private/bufpkg/bufplugin/bufpluginstore"
	"github.com/bufbuild/buf/private/bufpkg/bufregistryapi/bufregistryapimodule"
	"github.com/bufbuild/buf/private/bufpkg/bufregistryapi/bufregistryapiowner"
	"github.com/bufbuild/buf/private/bufpkg/bufregistryapi/bufregistryapiplugin"
	"github.com/bufbuild/buf/private/pkg/app/appext"
	"github.com/bufbuild/buf/private/pkg/filelock"
	"github.com/bufbuild/buf/private/pkg/normalpath"
	"github.com/bufbuild/buf/private/pkg/storage/storageos"
	"github.com/bufbuild/buf/private/pkg/wasm"
)

var (
	// AllCacheRelDirPaths are all directory paths for all time
	// concerning the module and plugin caches.
	//
	// These are normalized.
	// These are relative to container.CacheDirPath().
	//
	// This variable is used for clearing the cache.
	AllCacheRelDirPaths = []string{
		v1CacheModuleDataRelDirPath,
		v1CacheModuleLockRelDirPath,
		v1CacheModuleSumRelDirPath,
		v1beta1CacheModuleDataRelDirPath,
		v1beta1CacheModuleLockRelDirPath,
		v2CacheModuleRelDirPath,
		v3CacheCommitsRelDirPath,
		v3CacheModuleLockRelDirPath,
		v3CacheModuleRelDirPath,
		v3CachePluginRelDirPath,
		v3CacheWKTRelDirPath,
		v3CacheWasmRuntimeRelDirPath,
	}

	// v1CacheModuleDataRelDirPath is the relative path to the cache directory where module data
	// was stored in v1beta1.
	//
	// Normalized.
	v1beta1CacheModuleDataRelDirPath = "mod"
	// v1CacheModuleLockRelDirPath is the relative path to the cache directory where module lock files
	// were stored in v1beta1.
	//
	// Normalized.
	v1beta1CacheModuleLockRelDirPath = normalpath.Join("lock", "mod")
	// v1CacheModuleDataRelDirPath is the relative path to the cache directory where module data is stored.
	//
	// Normalized.
	// This is where the actual "clones" of the modules are located.
	v1CacheModuleDataRelDirPath = normalpath.Join("v1", "module", "data")
	// v1CacheModuleLockRelDirPath is the relative path to the cache directory where module lock files are stored.
	//
	// Normalized.
	// These lock files are used to make sure that multiple buf processes do not corrupt the cache.
	v1CacheModuleLockRelDirPath = normalpath.Join("v1", "module", "lock")
	// v1CacheModuleSumRelDirPath is the relative path to the cache directory where module digests are stored.
	//
	// Normalized.
	// These digests are used to make sure that the data written is actually what we expect, and if it is not,
	// we clear an entry from the cache, i.e. delete the relevant data directory.
	v1CacheModuleSumRelDirPath = normalpath.Join("v1", "module", "sum")
	// v2CacheModuleRelDirPath is the relative path to the cache directory for content addressable storage.
	//
	// Normalized.
	// This directory replaces the use of v1CacheModuleDataRelDirPath, v1CacheModuleLockRelDirPath, and
	// v1CacheModuleSumRelDirPath with a cache implementation using content addressable storage.
	v2CacheModuleRelDirPath = normalpath.Join("v2", "module")
	// v3CacheModuleRelDirPath is the relative path to the files cache directory in its newest iteration.
	//
	// Normalized.
	v3CacheModuleRelDirPath = normalpath.Join("v3", "modules")
	// v3CacheCommitsRelDirPath is the relative path to the commits cache directory in its newest iteration.
	//
	// Normalized.
	v3CacheCommitsRelDirPath = normalpath.Join("v3", "commits")
	// v3CacheWKTRelDirPath is the relative path to the well-known types cache directory in its newest iteration.
	//
	// Normalized.
	v3CacheWKTRelDirPath = normalpath.Join("v3", "wellknowntypes")
	// v3CacheModuleLockRelDirPath is the relative path to the lock files directory for module data.
	// This directory is used to store lock files for synchronizing reading and writing module data from the cache.
	//
	// Normalized.
	v3CacheModuleLockRelDirPath = normalpath.Join("v3", "modulelocks")
	// v3CachePluginRelDirPath is the relative path to the files cache directory in its newest iteration.
	//
	// Normalized.
	v3CachePluginRelDirPath = normalpath.Join("v3", "plugins")
	// v3CacheWasmRuntimeRelDirPath is the relative path to the Wasm runtime cache directory in its newest iteration.
	// This directory is used to store the Wasm runtime cache. This is an implementation specific cache and opaque outside of the runtime.
	//
	// Normalized.
	v3CacheWasmRuntimeRelDirPath = normalpath.Join("v3", "wasmruntime")
)

// NewModuleDataProvider returns a new ModuleDataProvider while creating the
// required cache directories.
func NewModuleDataProvider(container appext.Container) (bufmodule.ModuleDataProvider, error) {
	clientConfig, err := NewConnectClientConfig(container)
	if err != nil {
		return nil, err
	}
	return newModuleDataProvider(
		container,
		bufregistryapimodule.NewClientProvider(
			clientConfig,
		),
		bufregistryapiowner.NewClientProvider(
			clientConfig,
		),
	)
}

// NewCommitProvider returns a new CommitProvider while creating the
// required cache directories.
func NewCommitProvider(container appext.Container) (bufmodule.CommitProvider, error) {
	clientConfig, err := NewConnectClientConfig(container)
	if err != nil {
		return nil, err
	}
	return newCommitProvider(
		container,
		bufregistryapimodule.NewClientProvider(
			clientConfig,
		),
		bufregistryapiowner.NewClientProvider(
			clientConfig,
		),
	)
}

// NewPluginDataProvider returns a new PluginDataProvider while creating the
// required cache directories.
func NewPluginDataProvider(container appext.Container) (bufplugin.PluginDataProvider, error) {
	clientConfig, err := NewConnectClientConfig(container)
	if err != nil {
		return nil, err
	}
	return newPluginDataProvider(
		container,
		bufregistryapiplugin.NewClientProvider(
			clientConfig,
		),
	)
}

// NewWasmRuntime returns a new Wasm runtime while creating the required cache
// directories.
func NewWasmRuntime(ctx context.Context, container appext.Container) (wasm.Runtime, error) {
	// This is used by the Wasm runtime to cache compiled Wasm plugins. This is an
	// implementation specific cache and opaque outside of the runtime. The runtime
	// will manage the cache versioning itself within this directory.
	if err := createCacheDir(container.CacheDirPath(), v3CacheWasmRuntimeRelDirPath); err != nil {
		return nil, err
	}
	fullCacheDirPath := normalpath.Join(container.CacheDirPath(), v3CacheWasmRuntimeRelDirPath)
	wasmRuntime, err := wasm.NewRuntime(ctx, wasm.WithLocalCacheDir(fullCacheDirPath))
	if err != nil {
		return nil, err
	}
	return wasmRuntime, nil
}

// NewWKTStore returns a new bufwktstore.Store while creating the required cache directories.
func NewWKTStore(container appext.Container) (bufwktstore.Store, error) {
	if err := createCacheDir(container.CacheDirPath(), v3CacheWKTRelDirPath); err != nil {
		return nil, err
	}
	fullCacheDirPath := normalpath.Join(container.CacheDirPath(), v3CacheWKTRelDirPath)
	// No symlinks.
	storageosProvider := storageos.NewProvider()
	cacheBucket, err := storageosProvider.NewReadWriteBucket(fullCacheDirPath)
	if err != nil {
		return nil, err
	}
	return bufwktstore.NewStore(
		container.Logger(),
		cacheBucket,
	), nil
}

func newModuleDataProvider(
	container appext.Container,
	moduleClientProvider bufregistryapimodule.ClientProvider,
	ownerClientProvider bufregistryapiowner.ClientProvider,
) (bufmodule.ModuleDataProvider, error) {
	if err := createCacheDir(container.CacheDirPath(), v3CacheModuleRelDirPath); err != nil {
		return nil, err
	}
	fullCacheDirPath := normalpath.Join(container.CacheDirPath(), v3CacheModuleRelDirPath)
	delegateModuleDataProvider := bufmoduleapi.NewModuleDataProvider(
		container.Logger(),
		moduleClientProvider,
		newGraphProvider(container, moduleClientProvider, ownerClientProvider),
	)
	// No symlinks.
	storageosProvider := storageos.NewProvider()
	cacheBucket, err := storageosProvider.NewReadWriteBucket(fullCacheDirPath)
	if err != nil {
		return nil, err
	}
	if err := createCacheDir(container.CacheDirPath(), v3CacheModuleLockRelDirPath); err != nil {
		return nil, err
	}
	filelocker, err := filelock.NewLocker(normalpath.Join(container.CacheDirPath(), v3CacheModuleLockRelDirPath))
	if err != nil {
		return nil, err
	}
	return bufmodulecache.NewModuleDataProvider(
		container.Logger(),
		delegateModuleDataProvider,
		bufmodulestore.NewModuleDataStore(
			container.Logger(),
			cacheBucket,
			filelocker,
		),
	), nil
}

func newCommitProvider(
	container appext.Container,
	moduleClientProvider bufregistryapimodule.ClientProvider,
	ownerClientProvider bufregistryapiowner.ClientProvider,
) (bufmodule.CommitProvider, error) {
	if err := createCacheDir(container.CacheDirPath(), v3CacheCommitsRelDirPath); err != nil {
		return nil, err
	}
	fullCacheDirPath := normalpath.Join(container.CacheDirPath(), v3CacheCommitsRelDirPath)
	delegateReader := bufmoduleapi.NewCommitProvider(container.Logger(), moduleClientProvider, ownerClientProvider)
	// No symlinks.
	storageosProvider := storageos.NewProvider()
	cacheBucket, err := storageosProvider.NewReadWriteBucket(fullCacheDirPath)
	if err != nil {
		return nil, err
	}
	return bufmodulecache.NewCommitProvider(
		container.Logger(),
		delegateReader,
		bufmodulestore.NewCommitStore(
			container.Logger(),
			cacheBucket,
		),
	), nil
}

func newPluginDataProvider(
	container appext.Container,
	pluginClientProvider bufregistryapiplugin.ClientProvider,
) (bufplugin.PluginDataProvider, error) {
	if err := createCacheDir(container.CacheDirPath(), v3CachePluginRelDirPath); err != nil {
		return nil, err
	}
	fullCacheDirPath := normalpath.Join(container.CacheDirPath(), v3CachePluginRelDirPath)
	storageosProvider := storageos.NewProvider() // No symlinks.
	cacheBucket, err := storageosProvider.NewReadWriteBucket(fullCacheDirPath)
	if err != nil {
		return nil, err
	}
	delegateModuleDataProvider := bufpluginapi.NewPluginDataProvider(
		container.Logger(),
		pluginClientProvider,
	)
	return bufplugincache.NewPluginDataProvider(
		container.Logger(),
		delegateModuleDataProvider,
		bufpluginstore.NewPluginDataStore(
			container.Logger(),
			cacheBucket,
		),
	), nil
}

func createCacheDir(baseCacheDirPath string, relDirPath string) error {
	baseCacheDirPath = normalpath.Unnormalize(baseCacheDirPath)
	relDirPath = normalpath.Unnormalize(relDirPath)
	fullDirPath := filepath.Join(baseCacheDirPath, relDirPath)
	// OK to use os.Stat instead of os.LStat here as this is CLI-only
	fileInfo, err := os.Stat(fullDirPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return os.MkdirAll(fullDirPath, 0755)
		}
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf(
			"Expected %q to be a directory. This is used for buf's cache. "+
				"You can override the base cache directory %q by setting the $BUF_CACHE_DIR environment variable.",
			fullDirPath,
			baseCacheDirPath,
		)
	}
	if fileInfo.Mode().Perm()&0700 != 0700 {
		return fmt.Errorf(
			"Expected %q to be a writeable directory. This is used for buf's cache. "+
				"You can override the base cache directory %q by setting the $BUF_CACHE_DIR environment variable.",
			fullDirPath,
			baseCacheDirPath,
		)
	}
	return nil
}
