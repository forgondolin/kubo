package memory

import (
       "context"
       "errors"
       "io"

       filestore "github.com/ipfs/go-filestores"
       ds "github.com/ipfs/go-datastore"
       config "github.com/ipfs/kubo/config"

)

fx.Provide(func() datastore.Datastore {
  return ds.MutexWrap(datastore.NewMapDatastore())
})
