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
    image: ghcr.io/spejder/ms-vcard:latest
    restart: always
    ports:
      - '80/tcp'
    environment:
      MS_DATABASE: "dds"
      MS_URL: "https://medlem.dds.dk"
```

## vcard eksempel

Her er et eksempel på mine egne oplysninger genereret af servicen ud
fra data i Medlemsservice.

```vcard
BEGIN:VCARD
VERSION:4.0
ADR:;;Ejbysvinget 18;Glostrup;;2600;
BDAY:19730430
EMAIL:arne@arnested.dk
N:Jørgensen;Arne;;;
NICKNAME:Boga
NOTE:Medlemsnummer: 2014080
ORG:Ejby Gruppe
RELATED:urn:uuid:be90ae96-4c60-482c-b8fa-8dab6678dd60
TEL;TYPE=cell:+4521650113
UID:urn:uuid:741ecef2-866b-48fb-9460-695dc848c550
END:VCARD
```
