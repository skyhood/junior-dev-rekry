# Junior Software Developeriksi Duunitorille!

Moikka! Kiitos vielä hakemuksestasi ja onnea! Olemme valinneet sinut Junior Software Developer -rekrytointiprosessin seuraavaan vaiheeseen. Alta löydät ennakkotehtävän, jonka avulla selvitämme paremmin, sovimmeko yhteen. Aloitetaan yleisen tason infolla.

Voit valita vahvuuksiesi pohjalta haluatko tehdä frontend- *vai* backend-tehtävän. Valitse toinen tehtävistä, *molempia ei siis ole tarkoitus tehdä!*

Olemme suunnitelleet tehtävän kestämään *2–5 tuntia*, emmemekä odota, että käytät tähän aikaa sen enempää.

Emme odota sinun olevan valmis ammattilainen, eikä koodisikaan tarvitse olla täydellistä (eihän se koskaan ole). Etsimme kuitenkin tyyppiä, joka pystyy jo sopivasti tuettuna tekemään itsenäistä työtä. Tätä pyrimme tehtävällä tunnistamaan. Arviointikriteereistä voit lukea lisää tehtäväkuvausten alta.


## Tehtävä – Frontend

Löydät `jobs.json`-tiedostosta vinon pinon työpaikkailmoitusdataa (1 000 kpl, jos tarkkoja ollaan). Tee näistä työpaikkojen listaussivu.

Sivulla tulisi olla mahdollista nähdä olemassa olevat avoimet työpaikat. Sivulla voi olla erilaisia ominaisuuksia filttereistä animointeihin tai sivutuksesta upeaan ulkoasuun. Vain taivas (ja aika) on rajana.

Voit tehdä sivuston ilmeestä sellaisen kuin haluat. Inspiraatiota voit ottaa esimerkiksi [hakutulossivuiltamme](https://duunitori.fi/tyopaikat). Laitoimme liitteeksi myös css-tiedoston, josta löytyy hieman tukea tähän :) Visuaalinen ilme ei ole pakollinen arviointikriteeri, mutta odotamme toki sivun palvelevan tarkoitustaan.


### Arviointi
Yleiset arviointikriteerit löydät seuraavasta osiosta, tässä tämän tehtävän erityispiirteet.

Tässä tehtävässä voit valita, haluatko keskittyä työssäsi visualisuuteen, tehokkuuteen, saavutettavuuteen ja/vai käytettävyyteen. Valitse 1–2 näistä asioista ja kerro, mihin näistä olet keskittynyt. Näin osaamme arvioida tuotostasi oikeasta perspektiivistä :)


#### Visuaalisuus
Meillä devaajat ja designerit etsivät parhaita tapoja ratkaista kysymyksiä yhdessä. Kun devi ja design löytävät yhteisen sävelen, tuloksena on jotain upeaa.


#### Tehokkuus
Meillä, kuten myös tässä tehtävässä, on valtavasti dataa. Kaikki eivät kuitenkaan etsi töitä uusimmilla vempeleillä ja nopeimmilla nettiyhteyksillä, ja tämän pyrimme huomioimaan sivustomme kehityksessä.

#### Saavutettavuus
Me haluamme tehdä työelämästä monimuotoista, joten koodin saavutettavuus ja esteellisyyden poistaminen on yksi tavoitteistamme.


#### Käytettävyys
Ärsyttävät palvelut ovat ärsyttäviä. Emme halua lisätä ärsytystä, vaan tehdä palvelukokemustamme käytännöllisen, toimivan ja kenties jopa ilahduttavan.



## Tehtävä – Backend

Löydät `jobs.json`-tiedostosta vinon pinon työpaikkailmotusdataa (1 000 kpl, jos ihan tarkkoja ollaan). Tee näistä API. Rajapintaa käyttää ensisijaisesti nettisivu, joka listaa ja filtteröi avoimia työpaikkailmoituksia sekä näyttää ilmoitusten tietoja.

Voit hyödyntää vapaasti valitsemaasi koodikieltä ja apukirjastoja. Jos kaipaat inspiraatiota stäkin valintaan, meillä käytetään sekä Pythonia Djangon kera että Gota.

Sinun ei tarvitse luoda tuotantolaatuista palvelinta. Tee kuitenkin jonkinlainen kyselyitä vastaanottava palvelu, löydät pari esimerkkiä `backend_samples`-kansiosta.

### Ominaisuuksia

Tässä muutama esimerkkiominaisuus, joita voit implementoida. Kaikkea ei tarvitse tehdä, vaan valitse mitä haluat ja ehdit annetussa ajassa tekemään.

#### Haku
Primitiivisen hakutoiminnon implementointi. Haku on meillä Duunitorilla luonnollisesti palvelun ytimessä, sillä ihmiset ovat kiinnostuneita omista (ja naapurin) töistä.

#### Tilastodata
Datan aggregointi kiinnostaviksi luvuiksi tai dataseteiksi. Mitä, missä, milloin, miksi, miten paljon? Esimerkkidatan pohjalta mahdollisuudet ovat hieman rajalliset, mutta meillä on valtavasti erilaista dataa, ja tätä pyrimme hyödyntämään monin tavoin.

#### Sivutus
Implementoi rajapinnan palauttamiin tuloksiin jonkinlainen sivutus. Meiltä löytyy työpaikkailmoituksia joka hetki noin 50 000. Ihan jokaista niistä ei kannata kerralla silmille läväyttää.

#### Top-listat
Implementoi rajapinta erilaisille kiinnostaville top-listoille. Listat voivat olla jotain tällaista: ”7 kuuminta työnantajaa – tutustu näihin, jos haluat työllistyä”, ”Muutto edessä? Katso 12 kuntaa, joista löytyy töitä!” Meillä kahlataan jatkuvasti dataa, jotta voimme tuottaa kiinnostavia sisältöjä yleisöllemme ja medialle.

#### Yksittäinen ilmoitus
Implementoi rajapinta, joka palauttaa yksittäisen ilmoituksen sisällön. Vaikka listat ovat kivoja, vielä kivempaa on, kun sieltä löytää jotain itseään kiinnostavaa.


### Arviointi
Yleiset arviointikriteerit löydät seuraavasta osiosta, tässä tämän tehtävän erityispiirteet.

Kiinnitä tässä tehtävässä erityistä huomiota koodin tehokkuuteen. Huomioithan myös käyttötarkoituksen – tarvitseeko nettisaitti kerralla 1 000 työpaikkaa kaikkine tietoineen? Meillä, kuten myös tässä tehtävässä, on valtavasti dataa. Kaikki eivät kuitenkaan etsi töitä uusimmilla vempeleillä ja nopeimmilla nettiyhteyksillä, ja tämän pyrimme huomioimaan sivustomme kehityksessä.


## Arviointikriteerit

Pyrimme selkeyteen ja ymmärrettävyyteen arvioinnissamme. Tehtäväkohtaisten kuvausten lisäksi kiinnitämme erityistä huomiota seuraaviin asioihin:

### Koodin toimivuus
Homman tulee lähtökohtaisesti toimia. Fronttituotosta tutkimme lähtökohtaisesti uusimmalla Chromella. Muiden selaimien tukemisesta ei tarvitse huolehtia. Liitäthän mukaan selkeät ohjeet, joiden avulla saamme softasi pyörimään. Meillä ei valitettavasti ole aikaa puutteellisesti toimivan koodin debuggaamiseen, käytämme sen mieluummin oman koodimme bugien liiskaamiseen :)

### Koodin ymmärrettävyys
Meillä työskennellään porukalla samojen asioiden parissa, joten muidenkin on ymmärrettävä kirjoitettua koodia. Noudata hyviä käytäntöjä nimeämisten, kommenttien, dokumentoinnin ja muiden vastaavien suhteen. Olisi valitettavaa, jos koodisi toimisi upeasti, mutta emme pystyisi ymmärtämään, miksi.

### Koodin uudelleenkäytettävyys
Jotta meidän ei tarvitse kirjoittaa samanlaisia koodirivejä uudestaan ja uudestaan, haluamme tunnistaa uudelleenkäytettävät osat koodistamme. Tällaista perspektiiviä omaa tekemistä kohtaan arvostamme myös sinulta. [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself) on mukava periaate (joskaan sitäkään ei kannata ylisuorittaa, been there.)

### Versionhallinta
Meillä käytetään gittiä, ja osaaminen vähintään gitin perusteista katsotaan eduksi. Lopputuloksen toimittaminen ymmärrettävän git-historian kanssa on mahdollinen lisäpisteen paikka, mutta ei vaatimus.

## jobs.json

JSON-tiedosto näyttää tältä

```
{
  "jobs": [
    {
      "heading": string
      "date_posted": string
      "slug": string
      "municipality_name": string
      "export_image_url": string
      "company_name": string
      "descr": string
      "latitude": int
      "longitude": int
      "salary": {
        "type": int
        "min": int
        "max": int
      }
    }, ...
  ]
}
```


## Valmiin tuotoksen toimittaminen meille

Tadaa, valmista! Voit palauttaa tehtävän haluamallasi tavalla. Itse käytämme GitHubia, ja uusi yksityinen repository on oikein toimiva tapa toimittaa valmis lopputulos. Huomioithan kuitenkin muutaman asian.

1. *Älä forkkaa tätä repositoryä tai tee pull requestia suoraan tänne!* Muutoin työsi näkyy julkisesti kaikille, ja olisi ikävää, jos kanssahakija kopioisi työtäsi.
2. *Jos hyödynnät GitHubia tai vastaavaa, älä tee repositorystasi julkista* samasta syystä.
3. Toimitathan työsi mukana mahdollisimman selkeät ohjeet, joiden avulla pystymme kokeilemaan tuotostasi. Meillä ei valitettavasti ole aikaa ratkoa bugeja koodistasi :)


## Kysymyksiä?

Heräsikö ennakkotehtävästä tai työnhakuprosessista jotain kysyttävää? Kerron mielelläni lisää! Tavoitat minut parhaiten sähköpostilla osoitteesta [petri@duunitori.fi](mailto:petri@duunitori.fi) tai puhelimella numerosta [+358 440 422 294](tel:+358440422294).
