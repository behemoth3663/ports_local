# $FreeBSD$

PORTNAME=	gping
PORTVERSION=	0.1.6
DISTVERSIONPREFIX=	v
CATEGORIES=	net

MAINTAINER=	none@freebsd.org
COMMENT=	Ping with a graph

LICENSE=	MIT
#LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo
# ssl
USE_GITHUB=	yes
GH_ACCOUNT=	orf

#https://github.com/agrinman/tunnelto/releases/
#GH_TAGNAME=	ab321d9

JQ_GETVER_EXPR=	[.[].name | select(test("^[0-9.]+$$"; "n"))][0]

BROKEN=	'Error pinging: Unsupported OS Unknown, src/main.rs:122:33'

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#rm -f -- distinfo Makefile.deps && make makesum && make cargo-crates && | make cargo-crates >Makefile.deps && make makesum
.sinclude "${.CURDIR}/Makefile.deps"

PLIST_FILES=	bin/${PORTNAME}

#CARGO_INSTALL=
#do-install:
#.for _i in ${PLIST_FILES}
#	${INSTALL_PROGRAM} ${WRKDIR}/target/release/${_i:C|.*/||} ${STAGEDIR}${PREFIX}/${_i}
#.endfor

#post-install:
#.for _i in ${PLIST_FILES:Mbin/*}
##	${MV} ${STAGEDIR}${PREFIX}/${_i:S/${PORTNAME}/wormhole/} ${STAGEDIR}${PREFIX}/${_i}
#	${STRIP_CMD} ${STAGEDIR}${PREFIX}/${_i}
#.endfor

.include <bsd.port.mk>
