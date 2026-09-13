PORTNAME=	convert2json
DISTVERSION=	2.6.4
CATEGORIES=	local converters

MAINTAINER=	none@freebsd.org
COMMENT=	Convert BSON,CBOR,CSV,INI,MessagePack,Plist,RSV,TOML,XML,YAML to JSON
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE-MIT

RUN_DEPENDS=	jq:textproc/jq

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	simonrupf

CONFLICTS=	textproc/yq

PLIST_FILES=	\
	bin/bson2json \
	bin/bsonq \
	bin/cbor2json \
	bin/cborq \
	bin/cq \
	bin/csv2json \
	bin/ini2json \
	bin/iq \
	bin/msg2json \
	bin/msgq \
	bin/plist2json \
	bin/plistq \
	bin/rq \
	bin/rsv2json \
	bin/toml2json \
	bin/tq \
	bin/xml2json \
	bin/xq \
	bin/yaml2json \
	bin/yq

FIREFOX_VERSION_CMD=	${PKG_QUERY} --evaluate %n=firefox %v | ${SED} -r 's/^([0-9]+\.[0-9]+).*/\1/'
FETCH_ENV=	HTTP_USER_AGENT='Mozilla/5.0 (X11; Linux x86_64; rv:${FIREFOX_VERSION_CMD:sh}) Gecko/20100101 Firefox/${FIREFOX_VERSION_CMD:sh}'

post-install: .PHONY
.for _i in ${PLIST_FILES}
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/${_i}
.endfor

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${.CURDIR}/Makefile.crates )
	( ${MAKE} -C ${.CURDIR} FETCH_ENV="${FETCH_ENV}" makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
