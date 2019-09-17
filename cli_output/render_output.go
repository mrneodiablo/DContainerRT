package cli_output

import (
	"main/images"
	"text/tabwriter"
	"os"
	"fmt"
)	

func PrintListImage(listImage []images.ImageInfo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 3, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "IMAGES\tTAG\tID\tCREATED\t")

	for index := 0; index < len(listImage); index++ {
		fmt.Fprintln(w, listImage[index].Name + "\t" + listImage[index].Tag + "\t"+ listImage[index].Id + "\t" + string(listImage[index].Created) + "\t")
	}
	fmt.Fprintln(w)
	w.Flush()
}
