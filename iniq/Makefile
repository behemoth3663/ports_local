PORTNAME=	iniq
DISTVERSIONPREFIX=	v
DISTVERSION=	0.4.0
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Simple INI file reader for command line
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		gmake perl5
USE_PERL5=	build

USE_GITHUB=	yes
GH_ACCOUNT=	jcrd

MAKE_ARGS=	PREFIX=${PREFIX}

PLIST_FILES=	bin/${PORTNAME} \
		share/man/man1/${PORTNAME}.1.gz

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

.include <bsd.port.mk>
