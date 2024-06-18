package vmchecker

type VMCreateRequestBody struct {
	Name                  string
	FlavorID              string
	ExternalIP            string
	InternalIP            string
	SelectedOS            string
	UnionmountImage       string
	Keypair               string
	SelectedSecuritygroup string
	UserID                string
}
