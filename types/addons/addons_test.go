package addons

import "testing"
import "bytes"
import "encoding/json"

var data = []byte(`{
"id":2756696,
"user_id":33409,
"from_url":"postgres://xxx:xxx@ec2-50-17-200-129.compute-1.amazonaws.com/xxxxx",
"from_name":"SHARED_DATABASE (DATABASE_URL)",
"to_url":"s3://hkpgbackups/app2499171@heroku.com/b001.dump",
"to_name":"BACKUP",
"created_at":"2012/08/03 03:34.26",
"started_at":"2012/08/03 03:34.26",
"updated_at":"2012-08-03 03:34:31 +0000",
"finished_at":"2012/08/03 03:34.31",
"error_at":null,
"destroyed_at":null,
"progress":"upload done",
"size":"96.8KB",
"expire":true,
"cleaned_at":null,
"user":"app2499171@heroku.com",
"type":"backup",
"duration":4.873743,
"public_url":"https://s3.amazonaws.com/hkpgbackups/app2499171@heroku.com/b001.dump?AWSAccessKeyId=xxxx&Expires=1381576591&Signature=...."}`)

const id = 2756696
const userId = 33409
const fromUrl = "postgres://xxx:xxx@ec2-50-17-200-129.compute-1.amazonaws.com/xxxxx"
const fromName = "SHARED_DATABASE (DATABASE_URL)"
const toUrl = "s3://hkpgbackups/app2499171@heroku.com/b001.dump"
const toName = "BACKUP"
const createdAt = "2012/08/03 03:34.26"
const startedAt = "2012/08/03 03:34.26"
const updatedAt = "2012-08-03 03:34:31 +0000"
const finishedAt = "2012/08/03 03:34.31"
const errorAt = ""
const destroyedAt = ""
const progress = "upload done"
const size = "96.8KB"
const expire = true
const cleaned_at = ""
const user = "app2499171@heroku.com"
const aType = "backup"
const duration = 4.873743
const publicURl = "https://s3.amazonaws.com/hkpgbackups/app2499171@heroku.com/b001.dump?AWSAccessKeyId=xxxx&Expires=1381576591&Signature=...."

func TestPgBackupJSONDecoding(t *testing.T) {
	var backup *PgBackup
	reader := bytes.NewReader(data)
	err := json.NewDecoder(reader).Decode(&backup)
	if err != nil {
		t.Fatalf("decoding should succeed: %s", err.Error())
	}

	if backup.Id != id {
		t.Errorf("expected Id to be %d, was %d", id, backup.Id)
	}

	if backup.UserId != userId {
		t.Errorf("expected UserId to be %d, was %d", userId, backup.UserId)
	}

	if backup.FromUrl != fromUrl {
		t.Errorf("expected FromUrl to be `%s`, was `%s`", fromUrl, backup.FromUrl)
	}

	if backup.ToName != toName {
		t.Errorf("expected toName to be `%s`, was `%s`", toName, backup.ToName)
	}

	if backup.ErrorAt != errorAt {
		t.Errorf("expected ErrorAt to be `%s`, was `%s`", errorAt, backup.ErrorAt)
	}

	if backup.Expire != expire {
		t.Errorf("expected Expire to be %b, was %b", expire, backup.Expire)
	}

}

func TestPgBackupNameExists(t *testing.T) {
	backup := PgBackup{ToUrl: toUrl}
	name := backup.Name()
	if name != "b001" {
		t.Fatalf("expected `%s` to == `b001`", name)
	}
}

func TestPgBackupNameEmpty(t *testing.T) {
	backup := PgBackup{}
	name := backup.Name()
	if name != "" {
		t.Fatalf("expected name to be empty was: %q", name)
	}
}
