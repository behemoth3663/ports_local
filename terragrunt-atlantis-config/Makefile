PORTNAME=	terragrunt-atlantis-config
PORTVERSION=	1.19.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Generate Atlantis config for Terragrunt projects
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	transcend-io

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.VERSION=${PORTVERSION}'"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

_OPTIONS_extract=	600:pre-post-extract
pre-post-extract: .PHONY
	${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=FiloSottile GHT_PROJ=age GHT_PATH=filippo.io/age )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=mozilla GHT_PROJ=sops GHT_PATH=go.mozilla.org/sops/v3 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/Azure/go-autorest/(.*)|\1\3/\2|}

#-google:go-genproto:49dd2c1f3d0b:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:49dd2c1f3d0b:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
GH_TUPLE:=	${GH_TUPLE:N*/google.golang.org/genproto/*}

#-googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|}

GH_TUPLE:=	${GH_TUPLE:N*/cloud.google.com/go/*}
GH_TUPLE+=	googleapis:google-cloud-go:compute/metadata/v0.2.3:googleapis_google_cloud_go_compute_metadata

#-hashicorp:go-secure-stdlib:v0.1.2:hashicorp_go_secure_stdlib_mlock/vendor/github.com/hashicorp/go-secure-stdlib/mlock
#-hashicorp:go-secure-stdlib:v0.1.2:hashicorp_go_secure_stdlib_strutil/vendor/github.com/hashicorp/go-secure-stdlib/strutil
#+hashicorp:go-secure-stdlib:mlock/v0.1.2:hashicorp_go_secure_stdlib_mlock
#+hashicorp:go-secure-stdlib:strutil/v0.1.2:hashicorp_go_secure_stdlib_strutil
#-hashicorp:go-secure-stdlib:v0.1.3:hashicorp_go_secure_stdlib_parseutil/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
#+hashicorp:go-secure-stdlib:parseutil/v0.1.3:hashicorp_go_secure_stdlib_parseutil/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
GH_TUPLE:=	${GH_TUPLE:C,([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/hashicorp/go-secure-stdlib/(mlock|strutil),\1\3/\2,}
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/hashicorp/go-secure-stdlib)/(.*)|\1\3/\2|}

#-hashicorp:vault:v1.5.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault/api
#+hashicorp:vault:v1.5.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault
GH_TUPLE:=	${GH_TUPLE:S|/vault/api|/vault|}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>
