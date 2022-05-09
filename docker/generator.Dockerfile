FROM openapitools/openapi-generator-cli:latest

VOLUME . /local

CMD "generate \
        -g go-server \
        -i /local/petfeederapi.yaml \
        --global-property=generateAliasAsModel \
        --additional-properties=outputAsLibrary=true,onlyInterfaces=true,router=chi,sourceFolder=internal/api \
        --git-repo-id="petfeederapi" \
        --git-user-id="FoxFurry" \
        -o /local/."