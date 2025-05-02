PORTNAME=	tfsec
PORTVERSION=	1.28.14
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	nulldutra@proton.me
COMMENT=	Security scanner for Terraform
WWW=		https://aquasecurity.github.io/${PORTNAME}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aquasecurity

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
