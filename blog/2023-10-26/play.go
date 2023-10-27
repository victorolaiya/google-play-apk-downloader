package play

type Access_Token struct{}

type Acquire_Request struct{}

func (Acquire_Request) Do(doc string) error {
   return nil
}

type Checkin struct{}

type Delivery struct{}

type Delivery_Request struct{}

func (Delivery_Request) Do(doc string, vc uint64) (*Delivery, error) {
   return nil, nil
}

type Details struct{}

type Details_Request struct{}

func (Details_Request) Do(doc string) (*Details, error) {
   return nil, nil
}

type Device struct{}

type Raw_Checkin struct{}

func (Raw_Checkin) Checkin() (*Checkin, error) {
   return nil, nil
}

func (*Raw_Checkin) Do(Device) error {
   return nil
}

type Raw_Refresh_Token struct{}

func (*Raw_Refresh_Token) Do(oauth_token string) error {
   return nil
}

func (Raw_Refresh_Token) Token() (*Refresh_Token, error) {
   return nil, nil
}

type Refresh_Token struct{}

func (Refresh_Token) Do() (*Access_Token, error) {
   return nil, nil
}

type Sync_Request struct{}

func (Sync_Request) Do(Device) error {
   return nil
}

/////////////////////////

type File struct {
   Package_Name string
   Version_Code uint64
}
