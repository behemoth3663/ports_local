PORTNAME=	zizmor
DISTVERSIONPREFIX=	v
DISTVERSION=	1.30.1
CATEGORIES=	local devel

MAINTAINER=	none@FreeBSD.org
COMMENT=	Static analysis for GitHub Actions
WWW=		https://docs.zizmor.sh/ https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		cargo

USE_GITHUB=	yes
GH_ACCOUNT=	zizmorcore

PLIST_FILES=	bin/${PORTNAME}

CARGO_INSTALL_PATH=	crates/${PORTNAME}

# avoiding fetch: https://crates.io/api/v1/crates/http-cache-reqwest/1.0.0-alpha.9/download?dummy=/rust/crates/http-cache-reqwest-1.0.0-alpha.9.crate: Forbidden
FIREFOX_VERSION!=	pkg-static query --evaluate %n=firefox %v | sed -r 's/^([0-9]+\.[0-9]+).*/\1/'
FETCH_ENV=	HTTP_USER_AGENT='Mozilla/5.0 (X11; Linux x86_64; rv:${FIREFOX_VERSION}) Gecko/20100101 Firefox/${FIREFOX_VERSION}'

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${.CURDIR}/Makefile.crates )
	( ${MAKE} -C ${.CURDIR} FETCH_ENV="${FETCH_ENV}" makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.mk>
