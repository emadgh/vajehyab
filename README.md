# Vajeh Yab
###VajehYab.com Api Package for golang

Simply create an instance of `VajehYab` and `Search`

```golang
vy := vajehyab.VajehYab{Developer: "YourDeveloperName"}
vajeh, err := vy.Search(update.Message.Text)
if err != nil {
	panic(err)
	return err
}
fmt.Println(vajeh.Data.Text.ToString()+"\n\nمنبع: "+vajeh.Data.Source.ToString())
```