PORTNAME=	packer
PORTVERSION=	1.12.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	brad@facefault.org
COMMENT=	Tool for creating machine images
WWW=		https://packer.io/

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cpe go:modules
CPE_VENDOR=	hashicorp

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.GitCommit=$$(${GIT_COMMIT})'"

pre-deps: .PHONY
	if [ -d '${PATCHDIR}_' ]; then \
		${RM} -r ${PATCHDIR}; \
	else \
		${MV} ${PATCHDIR} ${PATCHDIR}_; \
	fi

pre-deps-all: .PHONY

post-deps-all: .PHONY
	test ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
