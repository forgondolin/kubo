package memory

import (
       "context"
       "errors"
       "io"

       filestore "github.com/ipfs/go-filestores"
       ds "github.com/ipfs/go-datastore"
       config "github.com/ipfs/kubo/config"

)

    type MemRepo struct {
    C *config.Config
    W  Datastore
}

func (m *MemRepo) RunInMemoryRepo() (*config.Config, error) {
     return &m.C{
     W:  fx.Provide(ds.MutexWrap(datastore.NewMapDatastore())), nil
     }
}
