# zone2json
This utility converts the records of a DNS (BIND) zone file to JSON/XML/YAML objects.

# Usage

Here is a sample of DNS records from some zone file:

```
almroot@x:~(master)$ cat /tmp/zone-se | head -n 50 | tail -n 10
0-0-0.se.		7200	IN	NSEC	0-0-1.se. NS DS RRSIG NSEC
0-0-0.se.		3600	IN	DS	12412 8 1 29560ECA044CB59C5B17304CAAC3DB74E0C1E110
0-0-0.se.		3600	IN	DS	12412 8 2 E857ED7FC925F98C2E926B9E2574FB80A3913AB9438A1C75B39630D5 5546736F
0-0-0.se.		86400	IN	NS	ns1.loopia.se.
0-0-0.se.		86400	IN	NS	ns2.loopia.se.
0-0-1.se.		3600	IN	RRSIG	DS 8 2 3600 20220125191311 20220113131050 30015 se. hnDOgNfD1GfDYB6ZsJGFQAEdShfMIjy80I666OMNLuCgT94idxFm8F43 bJhXLolVkr31q/yJI/NLWIxxVo9/rRhi2RJyhT7QGB8O92xYYIbC1IKc 4FucnCN9ral9Toayldpw64qCtjM2cW8ssbjTBCIVlGIE3Ecyt0uoFrac DhDteNei7BUSPCYFoGgNeFicJuXfIpXJIzbeSe880CDKnapVEwMm6erm WiwEc28wjfFjXUnBMSLWcfnjBC0bjWYA4TQkxlXMlgHR1L2jlzp2M/SC vG5fTeT0nBAQC6KMhYxLdXaQNwd7tka+P2NLfN7u1tg/ydrhS+fp0rDa dvT5Uw==
0-0-1.se.		7200	IN	RRSIG	NSEC 8 2 7200 20220126032018 20220113051050 30015 se. QXyfeekfwM0fU8IyjN3ZgWzOYP5+5WWaR7ur9wI3Brh4aYT2mVa5T8Rh OTmuSVw+AjUuQ1p1DoU3aojkjRvUBGCySxaCQl7QltvQjwmIUxQLguoa IVB0nymfftDIRmEOI8nMkE9rAlE3DuKR+3l/v/ByGhMjby+HDxDTV6FR t+eUYP+SJIl6dEbpNqWcgFh0VYk4IhLbYGhJie3AnvWAdj0jqAJIXyiv lUvT3PRYd0ywVFnZKAelk184SU312TPH+AVH1JOQgizJXYxSup88wdHz etc2MXEnTEAt+ZA4gdiRU8+FuDbMBRyTR62saXe6bMQGsGuWpiAnYwcm OResBA==
0-0-1.se.		7200	IN	NSEC	0-1.se. NS DS RRSIG NSEC
0-0-1.se.		3600	IN	DS	12412 8 1 6833DFCE0F6CAA9C9AF711F361D010C8C2FF3685
0-0-1.se.		3600	IN	DS	12412 8 2 47783E3806F62788EF4E4C69D1AFE48262BEC34872E8C400132107A7 6D442D82
```

Now the same data is converted to JSON to be used in conjunction with [jq](https://stedolan.github.io/jq/) or similar tooling:

```
almroot@x:~(master)$ cat /tmp/zone-se | head -n 50 | tail -n 10 | zone2json
{"Hdr":{"Name":"0-0-0.se.","Rrtype":47,"Class":1,"Ttl":7200,"Rdlength":0},"NextDomain":"0-0-1.se.","TypeBitMap":[2,43,46,47]}
{"Hdr":{"Name":"0-0-0.se.","Rrtype":43,"Class":1,"Ttl":3600,"Rdlength":0},"KeyTag":12412,"Algorithm":8,"DigestType":1,"Digest":"29560ECA044CB59C5B17304CAAC3DB74E0C1E110"}
{"Hdr":{"Name":"0-0-0.se.","Rrtype":43,"Class":1,"Ttl":3600,"Rdlength":0},"KeyTag":12412,"Algorithm":8,"DigestType":2,"Digest":"E857ED7FC925F98C2E926B9E2574FB80A3913AB9438A1C75B39630D55546736F"}
{"Hdr":{"Name":"0-0-0.se.","Rrtype":2,"Class":1,"Ttl":86400,"Rdlength":0},"Ns":"ns1.loopia.se."}
{"Hdr":{"Name":"0-0-0.se.","Rrtype":2,"Class":1,"Ttl":86400,"Rdlength":0},"Ns":"ns2.loopia.se."}
{"Hdr":{"Name":"0-0-1.se.","Rrtype":46,"Class":1,"Ttl":3600,"Rdlength":0},"TypeCovered":43,"Algorithm":8,"Labels":2,"OrigTtl":3600,"Expiration":1643137991,"Inception":1642079450,"KeyTag":30015,"SignerName":"se.","Signature":"hnDOgNfD1GfDYB6ZsJGFQAEdShfMIjy80I666OMNLuCgT94idxFm8F43bJhXLolVkr31q/yJI/NLWIxxVo9/rRhi2RJyhT7QGB8O92xYYIbC1IKc4FucnCN9ral9Toayldpw64qCtjM2cW8ssbjTBCIVlGIE3Ecyt0uoFracDhDteNei7BUSPCYFoGgNeFicJuXfIpXJIzbeSe880CDKnapVEwMm6ermWiwEc28wjfFjXUnBMSLWcfnjBC0bjWYA4TQkxlXMlgHR1L2jlzp2M/SCvG5fTeT0nBAQC6KMhYxLdXaQNwd7tka+P2NLfN7u1tg/ydrhS+fp0rDadvT5Uw=="}
{"Hdr":{"Name":"0-0-1.se.","Rrtype":46,"Class":1,"Ttl":7200,"Rdlength":0},"TypeCovered":47,"Algorithm":8,"Labels":2,"OrigTtl":7200,"Expiration":1643167218,"Inception":1642050650,"KeyTag":30015,"SignerName":"se.","Signature":"QXyfeekfwM0fU8IyjN3ZgWzOYP5+5WWaR7ur9wI3Brh4aYT2mVa5T8RhOTmuSVw+AjUuQ1p1DoU3aojkjRvUBGCySxaCQl7QltvQjwmIUxQLguoaIVB0nymfftDIRmEOI8nMkE9rAlE3DuKR+3l/v/ByGhMjby+HDxDTV6FRt+eUYP+SJIl6dEbpNqWcgFh0VYk4IhLbYGhJie3AnvWAdj0jqAJIXyivlUvT3PRYd0ywVFnZKAelk184SU312TPH+AVH1JOQgizJXYxSup88wdHzetc2MXEnTEAt+ZA4gdiRU8+FuDbMBRyTR62saXe6bMQGsGuWpiAnYwcmOResBA=="}
{"Hdr":{"Name":"0-0-1.se.","Rrtype":47,"Class":1,"Ttl":7200,"Rdlength":0},"NextDomain":"0-1.se.","TypeBitMap":[2,43,46,47]}
{"Hdr":{"Name":"0-0-1.se.","Rrtype":43,"Class":1,"Ttl":3600,"Rdlength":0},"KeyTag":12412,"Algorithm":8,"DigestType":1,"Digest":"6833DFCE0F6CAA9C9AF711F361D010C8C2FF3685"}
{"Hdr":{"Name":"0-0-1.se.","Rrtype":43,"Class":1,"Ttl":3600,"Rdlength":0},"KeyTag":12412,"Algorithm":8,"DigestType":2,"Digest":"47783E3806F62788EF4E4C69D1AFE48262BEC34872E8C400132107A76D442D82"}
```

# Features

Output from `--help`:

```
almroot@x:~$ zone2json --help
Usage:
  zone2json [OPTIONS]

Application Options:
  -i, --input=          The file to read from, use - for STDIN (default: -)
  -o, --output=         The file to write to, use - for STDOUT (default: -)
  -f, --format=         Output format: json/yaml/xml/bind (default: json)
      --origin=         The default origin for relative domains (default: .)
      --ttl=            The default TTL to be used (default: 86400)
      --allow-includes  Enables support for bind $INCLUDE directives

Help Options:
  -h, --help            Show this help message
```

# Building

This will produce a binary `zone2json` in the current working directory:

`almroot@x:~(main)$ go build -o zone2json ./main`
