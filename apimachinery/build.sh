# /bin/bash
rm -rf client
rm -rf models
rm -rf restapi/operations
swagger generate server -P models.Principal -f ./swagger.yaml --exclude-main --skip-models --existing-models=github.com/cloustone/pandas/models
swagger generate client -P models.Principal -f ./swagger.yaml --skip-models --existing-models=github.com/cloustone/pandas/models
