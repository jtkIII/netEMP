# netEMP

### request → normalize → event → filter → action

#### So Far
- main.go = wiring/bootstrap
- server package = HTTP concerns

#### Paradigm
- main owns startup
- Server owns routing
- http.Server owns sockets/connections
- handlers own request logic

#### Flow
HTTP request
    ↓
convert to Event
    ↓
process Event

###### jtkIII