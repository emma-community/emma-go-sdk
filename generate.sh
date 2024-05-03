openapi-generator generate \
  -i https://api.emma.ms/v3/api-docs \
  -g go \
  --additional-properties packageName=emma,packageVersion=0.0.1,useTags=true \
  --git-user-id emma-community \
  --git-repo-id emma-go-sdk