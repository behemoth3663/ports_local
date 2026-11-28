PORTNAME=	bat
DISTVERSIONPREFIX=	v
DISTVERSION=	0.27.0
CATEGORIES=	textproc
PKGNAMESUFFIX=	-local

MAINTAINER=	pizzamig@FreeBSD.org
COMMENT=	Clone of cat with syntax highlighting
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

LIB_DEPENDS=	libgit2.so:devel/libgit2 \
		libonig.so:devel/oniguruma

USES=		cargo cpe
CPE_VENDOR=	${PORTNAME}_project
USE_GITHUB=	yes
GH_ACCOUNT=	sharkdp

#share/zsh/site-functions/_bat
PLIST_FILES=	bin/bat \
		share/man/man1/bat.1.gz

post-install:
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/bat
	${FIND} ${WRKDIR}/target -name "bat.1" -exec ${CP} {} ${WRKSRC}/doc/bat.1 \;
#	${FIND} ${WRKDIR}/target -name "bat.zsh" -exec ${CP} {} ${WRKSRC}/doc/bat.zsh \;
	${INSTALL_MAN} ${WRKSRC}/doc/bat.1 ${STAGEDIR}${PREFIX}/share/man/man1
#	${MKDIR} ${STAGEDIR}${PREFIX}/share/zsh/site-functions
#	${INSTALL_DATA} ${WRKSRC}/doc/bat.zsh ${STAGEDIR}${PREFIX}/share/zsh/site-functions/_bat

FETCH_ENV=	HTTP_USER_AGENT='Mozilla/5.0 (X11; Linux x86_64; rv:${FIREFOX_VERSION_CMD:sh}) Gecko/20100101 Firefox/${FIREFOX_VERSION_CMD:sh}'

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
deps: .PHONY
	${RM} Makefile.crates
	( ${MAKE} -C ${.CURDIR} makesum )
	( set -o pipefail && ${MAKE} -C ${.CURDIR} cargo-crates | ${GREP} --extended-regexp --invert-match '^=' > ${.CURDIR}/Makefile.crates )
	( ${MAKE} -C ${.CURDIR} FETCH_ENV="${FETCH_ENV}" makesum )

.sinclude "${.CURDIR}/Makefile.crates"
.include <bsd.port.pre.mk>

FIREFOX_VERSION_CMD=	${PKG_QUERY} --evaluate %n=firefox %v | ${SED} -r 's/^([0-9]+\.[0-9]+).*/\1/'

.include <bsd.port.post.mk>
