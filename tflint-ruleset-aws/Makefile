PORTNAME=	tflint-ruleset-aws
PORTVERSION=	0.34.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	TFLint ruleset for terraform-provider-aws
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	terraform-linters

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-google:go-genproto:6ceb2ff114de:google_go_genproto/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:6ceb2ff114de:google_go_genproto/vendor/google.golang.org/genproto
GH_TUPLE:=	${GH_TUPLE:S|/googleapis/rpc$||}

.include <bsd.port.mk>
