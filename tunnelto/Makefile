#

PORTNAME=	tunnelto
#DISTVERSIONPREFIX=	v
PORTVERSION=	0.1.14
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Expose your local web server to the Internet via a public URL

LICENSE=	MIT
#LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo ssl
USE_GITHUB=	yes
GH_ACCOUNT=	agrinman

#https://github.com/agrinman/tunnelto/releases/
#GH_TAGNAME=	ab321d9

JQ_GETVER_EXPR=	[.[].name | select(test("^[0-9.]+$$"; "n"))][0]

fetch-version: .PHONY get-version

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#rm -f -- distinfo Makefile.deps && make makesum && make cargo-crates | grep -Ev '^=' >Makefile.deps && make makesum
.-include "${.CURDIR}/Makefile.deps"

PLIST_FILES=	bin/${PORTNAME} \
		bin/${PORTNAME}_server

CARGO_INSTALL=
do-install:
.for _i in ${PLIST_FILES}
	${INSTALL_PROGRAM} ${WRKDIR}/target/release/${_i:C|.*/||} ${STAGEDIR}${PREFIX}/${_i}
.endfor

#post-install:
#.for _i in ${PLIST_FILES:Mbin/*}
##	${MV} ${STAGEDIR}${PREFIX}/${_i:S/${PORTNAME}/wormhole/} ${STAGEDIR}${PREFIX}/${_i}
#	${STRIP_CMD} ${STAGEDIR}${PREFIX}/${_i}
#.endfor

.include <bsd.port.mk>
