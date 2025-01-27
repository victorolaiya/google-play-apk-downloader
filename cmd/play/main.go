package main

import (
   "154.pages.dev/google/play"
   "154.pages.dev/http"
   "flag"
   "fmt"
   "os"
   "strings"
)

type flags struct {
   acquire bool
   code string
   device bool
   doc string
   platform int64
   single bool
   version uint64
   trace bool
}

func main() {
   var f flags
   {
      var b strings.Builder
      b.WriteString("oauth_token from ")
      b.WriteString("accounts.google.com/embedded/setup/android")
      flag.StringVar(&f.code, "c", "", b.String())
   }
   flag.StringVar(&f.doc, "d", "", "doc")
   flag.BoolVar(&f.device, "device", false, "create device")
   flag.Int64Var(&f.platform, "p", 0, fmt.Sprint(play.Platforms))
   flag.BoolVar(&f.single, "s", false, "single APK")
   flag.BoolVar(&f.trace, "t", false, "trace")
   flag.Uint64Var(&f.version, "v", 0, "version code")
   flag.BoolVar(&f.acquire, "a", false, "acquire")
   flag.Parse()
   dir, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   dir += "/google/play"
   os.MkdirAll(dir, os.ModePerm)
   http.No_Location()
   if f.trace {
      http.Trace()
   } else {
      http.Verbose()
   }
   if f.code != "" {
      err := f.do_auth(dir)
      if err != nil {
         panic(err)
      }
   } else {
      play.Phone.Platform = play.Platforms[f.platform]
      switch {
      case f.device:
         err := f.do_device(dir)
         if err != nil {
            panic(err)
         }
      case f.doc != "":
         head, err := f.do_header(dir)
         if err != nil {
            panic(err)
         }
         switch {
         case f.acquire:
            err := head.Acquire(f.doc)
            if err != nil {
               panic(err)
            }
         case f.version >= 1:
            err := f.do_delivery(head)
            if err != nil {
               panic(err)
            }
         default:
            detail, err := head.Details(f.doc)
            if err != nil {
               panic(err)
            }
            fmt.Println(detail)
         }
      default:
         flag.Usage()
      }
   }
}
