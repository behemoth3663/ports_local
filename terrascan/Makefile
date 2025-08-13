PORTNAME=	terrascan
PORTVERSION=	1.18.12
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Compliance and security violations in terraform code
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
USE_GITLAB=	nodefault

GH_ACCOUNT=	accurics

PLIST_FILES=	bin/${PORTNAME}

#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

_OPTIONS_extract=	750:post-post-extract
post-post-extract: .PHONY
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api/api/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api || :
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml/kyaml/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml || :

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
