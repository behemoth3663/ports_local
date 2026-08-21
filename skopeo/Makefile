PORTNAME=	skopeo
PORTVERSION=	1.24.1
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX=	-local

MAINTAINER=	dfr@FreeBSD.org
COMMENT=	Inspect container images and repositories on registries
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

BUILD_DEPENDS+=	bash:shells/bash
BUILD_DEPENDS+=	curl:ftp/curl
BUILD_DEPENDS+=	go-md2man:textproc/go-md2man
LIB_DEPENDS=	libgpgme.so:security/gpgme
#RUN_DEPENDS=	${LOCALBASE}/etc/containers/containers.conf.sample:sysutils/containers-common

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

GO_MODULE=	github.com/containers/${PORTNAME}
GO_TARGET=	./cmd/${PORTNAME}

#CGO_ENABLED=0 -> imports github.com/proglottis/gpgme: build constraints exclude all Go files
#GO_LDFLAGS+=-extldflags -static -> ld: error: unable to find library -lgpgme
GO_BUILDFLAGS+=	-tags exclude_graphdriver_btrfs
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X go.podman.io/image/v5/signature/internal/sequoia.sequoiaLibraryDir=
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

pre-build: .PHONY
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io

post-stage: .PHONY
	${MKDIR} ${STAGEDIR}${PREFIX}/etc/containers/registries.d
	${INSTALL_DATA} ${WRKSRC}/default.yaml ${STAGEDIR}${PREFIX}/etc/containers/registries.d
	${FIND} ${WRKSRC}/docs -name '*.1.md' -type f | while read -r md; do \
		man=$${md%.md}; \
		${SED} -e 's/\((skopeo.*\.md)\)//i' -e 's/\[\(skopeo.*\)\]/\1/i' $${md} | go-md2man | \
			${SED} -r 's/^(\.TH)[[:blank:]]+skopeo([^[:blank:]]*)\(1\)(.*)/\1 skopeo\2 1\3/i' > $${man}; \
		${INSTALL_MAN} $${man} ${STAGEDIR}${PREFIX}/share/man/man1; \
	done

.include <bsd.port.mk>
