<p align="center">
  <a href="https://uselumen.co">
    <img src="https://user-images.githubusercontent.com/43097772/178112983-d1f040da-6580-473f-b1cc-6083a0c0c95e.png" height="60">
  </a>
  <p align="center">Data-driven automation messaging for growth and retention.</p>
</p>

# Lumen

Lumen-go client allows you to seamlessly identify and track user attributes and events on your app. Plus other perks.

## Features

- Identify users
- Track user events
- Update user properties

## Getting started

- Setup your [Lumen](https://uselumen.co) account.

- Retrieve your API key.

> Follow the steps below to retrieve your api key.
>
> - Log in to your lumen dashboard.
> - Navigate to Settings
> - Select the API Key tab to view and copy your key.

- Install the package

```sh
go get github.com/uselumen/lumen-go
```

## Installation

#### Import the package

```go
import (
 "github.com/uselumen/lumen-go"
)


```

#### Initialize the package

```go

const LumenApiKey = "<< your-api-key >>";

lumen := NewLumengo(LumenApiKey)

```

## Usage

#### Identify a user

```go

	data := IdentifyParams{
		Email:     "john@doe.co", // required
		FirstName: "Gopher",
		LastName:  "Basit",
	}

    err := lumen.Identify("<< user-identifier >>", data)

	if err != nil {
		return err
	}


```

#### Track an event

After identifying users, you can now capture their actions like "Product Clicked" or "Product Viewed" with other custom properties.

```go

 	properties := map[string]interface{}{
		"productId": 100023449,
	}


	err := lumen.Track("<< user-identifier >>", "<< event-name >>", params)

	if err != nil {
		return err
	}

```

## Contributing

1. Fork it
2. Clone your fork (`git clone git@github.com:MY_USERNAME/lumen-go.git && cd lumen-go`)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'feat: Added some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create new Pull Request
