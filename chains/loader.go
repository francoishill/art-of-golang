package main

type Loader func() *Error

func ChainLoad(loaders ...Loader) *Error {
	for _, l := range loaders {
		if err := l(); err != nil {
			return err
		}
	}
	return nil
}
