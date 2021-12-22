#

PORTNAME=	tunnelto
PORTVERSION=	0.1.18
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Expose your local web server to the Internet via a public URL

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo ssl
USE_GITHUB=	yes
GH_ACCOUNT=	agrinman

JQ_GETVER_EXPR=	[.[].name | select(test("^[0-9.]+$$"; "n"))][0]

fetch-version: .PHONY get-version

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#rm -f -- distinfo Makefile.deps && make makesum && make cargo-crates | grep -Ev '^=' >Makefile.deps && make makesum

PLIST_FILES=	bin/${PORTNAME} \
		bin/${PORTNAME}_server

CARGO_INSTALL_PATH=	${PLIST_FILES:Mbin/*:S|^bin|.|}

post-install: .PHONY
.for _i in ${PLIST_FILES:Mbin/*}
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/${_i}
.endfor

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
