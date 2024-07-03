PORTNAME=	packer
PORTVERSION=	1.11.1
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

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.GitCommit=$$(${GIT_COMMIT})'"

pre-deps: .PHONY
	if [ -d '${MASTERDIR}/files_' ]; then \
		${RM} -r ${MASTERDIR}/files; \
	else \
		${MV} ${MASTERDIR}/files ${MASTERDIR}/files_; \
	fi

pre-deps-all: .PHONY

post-deps-all: .PHONY
	test ! -d '${MASTERDIR}/files_' || ${MV} ${MASTERDIR}/files_ ${MASTERDIR}/files

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
