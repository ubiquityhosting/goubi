
# GoUbi

GoUbi is a Go client library for accessing the Ubiquity Hosting V2 API.

## Documentation

You can view the client API documentation at [http://godoc.org/github.com/ubiquityhosting/GoUbi](http://godoc.org/github.com/ubiquityhosting/GoUbi)

You can view Ubiquity Hosting API documentation at [http://api.ubiquityhosting.com/](http://api.ubiquityhosting.com/)

## Usage

First, import the Ubiquity Hosting API package into your project:

```go
import "github.com/ubiquityhosting/GoUbi"
```

Now, you are ready to create a new Ubiquity Hosting client instance:

```go
client := goubi.NewUbiClient(
	12345,									// your client ID 
	"ubic-12345", 							// your API user name
	"550d0457f1824ca0c127da6fdb8499a0"		// your API token
)
```
Once the client instance is created, it can then be used to access different services of the Ubiquity Hosting API.

## Functional Example

The following functional example shows how to get detailed information for images, flavors, and zones in order to create a new Ubiquity Hosting instance:

```go
client := goubi.NewUbiClient(
	12345,									// your client ID
	"ubic-12345", 							// your API user name
	"550d0457f1824ca0c127da6fdb8499a0"		// your API token
)

images, _ := client.Cloud.ListImages()

fmt.Println("\nImages:")
for _, image := range images {
	fmt.Printf("\tName: %-35s ID: %d\n", image.Name, image.Id)
}

flavors, _ := client.Cloud.ListFlavors()

fmt.Println("\nFlavors")
for _, flavor := range flavors {
	fmt.Printf("\tName: %-50s ID: %d\n", flavor.Name, flavor.Id)
}

zones, _ := client.Cloud.ListZones()

fmt.Println("\nZones:")
for _, zone := range zones {
	fmt.Printf("\tName: %s, ID: %d\n", zone.Name, zone.Id)
}

new_cloud := new(goubi.CreateVMParams)

new_cloud.Hostname = "My-Ubi-Instance"
new_cloud.ImageID = 18
new_cloud.FlavorID = 1
new_cloud.ZoneID = 7

fmt.Print("\nNow ordering a new Ubiquity Hosting instance...")
	
order_details, err := client.Cloud.Create(new_cloud)

if err != nil {
	panic(fmt.Sprintf("\nError: %s", err))
}

fmt.Printf("completed. Your cloud order ID is %d.\n\n", order_details.OrderID)
```

## Versioning

Each version of the client is tagged and the version is updated accordingly.

Since Go does not have a built-in versioning, a package management tool is
recommended - a good one that works with git tags is
[gopkg.in](http://labix.org/gopkg.in).

To see the list of past versions, run `git tag`.

## Authors

By [Justin Canington](mailto:justin.canington@nobistech.net) and [Andrew Ayers](mailto:andrew.ayers@nobistech.net) for Ubiquity Hosting.
