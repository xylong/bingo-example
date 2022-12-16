### 图书
```json
{
  "mappings": {
    "properties": {
      "id":    { "type": "integer" },
      "name":    {"type": "text","analyzer": "ik_max_word","search_analyzer": "ik_smart"},
      "blurb":  { "type": "text","analyzer": "ik_smart"  },
      "price1":   { "type": "float"},
      "price2":   { "type": "float"},
      "author":   { "type": "keyword"},
      "date":   { "type": "date","format": "yyyy-MM-dd","ignore_malformed": true},
      "kind":   { "type": "integer"},
      "press": { "type": "keyword"}
    }
  }
}
```
- PUT /books
- GET /books/_mapping
- GET /books/_count
- GET /books/_search
- GET /_cat/plugins