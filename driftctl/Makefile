PORTNAME=	driftctl
PORTVERSION=	0.40.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Detect, track and alert on infrastructure drift
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE.md

USES=		go:modules

#GO_MODULE=	github.com/snyk/${PORTNAME}
#GH_ACCOUNT=	${GO_MODULE:H:T}
#GH_PROJECT=	${GO_MODULE:T}
USE_GITHUB=	yes
GH_ACCOUNT=	snyk
GH_PROJECT=	${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/pkg/version.version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/build.env=release' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/build.enableUsageReporting=false'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
