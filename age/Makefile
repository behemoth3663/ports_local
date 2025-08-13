PORTNAME=	age
DISTVERSIONPREFIX=	v
DISTVERSION=	1.2.1
#PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	dmgk@FreeBSD.org
COMMENT=	Simple, modern and secure file encryption tool
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		go:modules

#GO_MODULE=	filippo.io/age
GO_TARGET=	./cmd/age \
		./cmd/age-keygen

USE_GITHUB=	yes
GH_ACCOUNT=	FiloSottile

GO_BUILDFLAGS+=	-ldflags="-s -w  -extldflags -static \
	-X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/age \
		bin/age-keygen \
		share/man/man1/age-keygen.1.gz \
		share/man/man1/age.1.gz

post-install: .PHONY
	${INSTALL_MAN} ${WRKSRC}/doc/*.1 ${STAGEDIR}${PREFIX}/share/man/man1/

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
