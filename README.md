# zerotier

a command line client tool for interacting with ZeroTier Central API

## Admin REST API

`https://my.zerotier.com/help/api`

### Supported* Endpoints

- [x] GET /api/status - Obtain information about this server and/or useful to the Central web UI.
- [-] GET/api/self - Get the currently authenticated user’s user record.
- [x] GET /api/randomToken - This generates a random token suitable for use as an API token server-side using a secure random source.
- [-] POST /api/auth/_logout - Hitting this endpoint causes the user to be logged out.
- [ ] POST /api/user/{userId} - Only fields marked as [rw] can be directly modified.
- [ ] GET /api/user/{userId} - get user details
- [ ] DELETE /api/user/{userId} - Delete a user. This will immediately and PERMANENTLY delete a user and all associated networks and data.
- [x] GET /api/network - Get all networks for which you have at least read access.
- [ ] POST /api/network/{networkId} - Only fields marked as [rw] can be directly modified.
- [x] GET /api/network/{networkId} - Get network details
- [x] DELETE /api/network/{networkId} - Delete a network and all its related information permanently.
- [x] GET/api/network/{networkId}/member - Get all members of a network for which you have at least read access.
- [x] GET /api/network/{networkId}/member/{nodeId} - Get network node details
- [ ] POST/api/network/{networkId}/member/{nodeId} - New members can be added to a network by POSTing them.

## environment config

create a config file in your home directory `~/.config/.zerotier.json`

```json
{
    "default": {
        "Host": "https://my.zerotier.com",
        "Key": "your api token"
    },
}
```

for support of multiple configurations create additional top level elements in your config file and pass a matching `--env` flag with your commands

if no `--env` flag is passed the cli will load the `default` configuration

## load build harness

```sh
make init
```

## load dependencies

```sh
GO111MODULE=on go mod vendor
```

## testing

- go run main.go service randomtoken | jq > schema/randomtoken.json
- go run main.go network get -n network_id | jq > schema/network.json
- go run main.go member list -n network_id | jq > schema/member_list.json
- go run main.go member get -n network_id -m member_id | jq > schema/member_list.json
- go run main.go member list -n network_id | jq > schema/member_list.json
- go run main.go member get -n network_id -m member_id | jq > schema/member.json

## notes

- GET /api/status - should include current auth user data but returns null
- GET /api/self - documented but not functional
- POST /api/auth/_logout - unsure the usefulness of this
