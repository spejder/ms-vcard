# Vcards for dine kontakter i Medlemsservice

Dette projekt udstiller dine kontakter i Medlemsservice som vcards på
et HTTP-endpoint (du bør altid have noget HTTPS foran).

Du skal bruge brugernavn og password til Medlemsservice som basic
authentication.

Jeg bruger
[ContactSync](https://play.google.com/store/apps/details?id=com.vcard.android)
på min Android mobil til at synkronisere mine kontakter i
Medlemsservice ved hjælp af dette projekt.

**Bemærk!** Opslaget i Medlemsservice er langsomt. Når den henter ~60
kontakter tager det cirka 6 sekunder og når den henter ~890 kontakter
tager det over et halvt minut. Så synkroniser ikke for tit og servicen
skalerer derfor nok ikke til mange brugere.

## docker-compose eksempel

```yaml
version: "2"

services:
  web:
    build: spejder/ms-vcard:latest
    restart: always
    ports:
      - '80/tcp'
    environment:
      MS_DATABASE: "dds"
      MS_URL: "https://medlem.dds.dk"
```
