package api

import (
        "encoding/json"
        "fmt"

        log "github.com/golang/glog"
        c "github.com/opensds/opensds/pkg/context"
        "github.com/opensds/opensds/pkg/controller/client"
        "github.com/opensds/opensds/pkg/db"
        "github.com/opensds/opensds/pkg/model"
)

func NewFileSharePortal() *FileSharePortal {
        return &FileSharePortal{
                CtrClient: client.NewClient(),
        }
}

type FileSharePortal struct {
        BasePortal

        CtrClient client.Client
}


func (v *FileSharePortal) CreateFileShare() {
        fmt.Println("I m in /pkg/api/fileshare.go")
        //if !policy.Authorize(v.Ctx, "fileshare:create") {
        //      return
        //}
        //ctx := c.GetContext(v.Ctx)
        var fileshare = model.FileShareSpec{
                BaseModel: &model.BaseModel{},
        }

        // Unmarshal the request body
        if err := json.NewDecoder(v.Ctx.Request.Body).Decode(&fileshare); err != nil {
              reason := fmt.Sprintf("Parse fileshare request body failed: %s", err.Error())
              v.Ctx.Output.SetStatus(model.ErrorBadRequest)
              v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
              log.Error(reason)
              return
        }
        result, err := CreateFileShareDBEntry(c.GetContext(v.Ctx), &fileshare)
        if err != nil {
                reason := fmt.Sprintf("Create fileshare failed: %s", err.Error())
                v.Ctx.Output.SetStatus(model.ErrorBadRequest)
                v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
                log.Error(reason)
                return
        }
        // Marshal the result.
        body, err := json.Marshal(result)
        if err != nil {
                reason := fmt.Sprintf("Marshal fileshare created result failed: %s", err.Error())
                v.Ctx.Output.SetStatus(model.ErrorBadRequest)
                v.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
                log.Error(reason)
                return
        }

        v.Ctx.Output.SetStatus(StatusAccepted)
        v.Ctx.Output.Body(body)
        return
}

func (v *FileSharePortal) ListFileShares() {
        //if !policy.Authorize(v.Ctx, "fileshare:list") {
        //        return
        //}
        m, err := v.GetParameters()
        if err != nil {
                errMsg := fmt.Sprintf("list fileshares failed: %s", err.Error())
                v.ErrorHandle(model.ErrorBadRequest, errMsg)
                return
        }
        result, err := db.C.ListFileSharesWithFilter(c.GetContext(v.Ctx), m)
        if err != nil {
                errMsg := fmt.Sprintf("list fileshares failed: %s", err.Error())
                v.ErrorHandle(model.ErrorInternalServer, errMsg)
                return
        }
        // Marshal the result.
        body, _ := json.Marshal(result)
        v.SuccessHandle(StatusOK, body)

        return
}
