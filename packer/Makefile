PORTNAME=	packer
PORTVERSION=	1.14.2
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	brad@facefault.org
COMMENT=	Tool for creating machine images
WWW=		https://packer.io/

LICENSE=	MPL20

USES=		cpe go:modules
CPE_VENDOR=	hashicorp

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.GitCommit=$$(${GIT_COMMIT})'"

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/api/transport/http/internal/propagation
	@${RM} -r ${WRKSRC}/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
	@test -z ${WRKSRC_hashicorp_go_secure_stdlib_parseutil} || ${RLN} ${WRKSRC_hashicorp_go_secure_stdlib_parseutil}/parseutil ${WRKSRC}/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
	@${MKDIR} ${WRKSRC}/vendor/github.com/hashicorp/go-getter
	@test -z ${WRKSRC_hashicorp_go_getter_v2} || ${RLN} ${WRKSRC_hashicorp_go_getter_v2} ${WRKSRC}/vendor/github.com/hashicorp/go-getter/v2
	@${MKDIR} ${WRKSRC}/vendor/github.com/hashicorp/go-getter/gcs
	@test -z ${WRKSRC_hashicorp_go_getter_gcs} || ${RLN} ${WRKSRC_hashicorp_go_getter_gcs}/gcs ${WRKSRC}/vendor/github.com/hashicorp/go-getter/gcs/v2
	@${MKDIR} ${WRKSRC}/vendor/github.com/hashicorp/go-getter/s3
	@test -z ${WRKSRC_hashicorp_go_getter_s3} || ${RLN} ${WRKSRC_hashicorp_go_getter_s3}/s3 ${WRKSRC}/vendor/github.com/hashicorp/go-getter/s3/v2
#	@${MKDIR} ${WRKSRC}/vendor/google.golang.org/genproto
#	@test -z ${WRKSRC_google_go_genproto} || ${RLN} ${WRKSRC_google_go_genproto}/googleapis ${WRKSRC}/vendor/google.golang.org/genproto
#	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/api
#	@test -z ${WRKSRC_google_go_genproto_api} || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/rpc
	@test -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=darccio GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=google-cloud-go GHT_PATH=cloud.google.com/go/compute/metadata )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-census-instrumentation:opencensus-proto:v0.4.1:census_instrumentation_opencensus_proto/vendor/github.com/census-instrumentation/opencensus-proto
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:*}

#-googleapis:google-cloud-go:v1.1.3:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.35.1:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
GH_TUPLE:=	${GH_TUPLE:N*/cloud.google.com/go/*}

#-hashicorp:consul:    v1.25.1:hashicorp_consul_api/vendor/github.com/hashicorp/consul/api
#+hashicorp:consul:api/v1.25.1:hashicorp_consul_api/vendor/github.com/hashicorp/consul
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/hashicorp/consul)/(.*)|\1\3/\2|}

#-hashicorp:go-secure-stdlib:        v0.1.2:hashicorp_go_secure_stdlib_strutil/vendor/github.com/hashicorp/go-secure-stdlib/strutil
#+hashicorp:go-secure-stdlib:strutil/v0.1.2:hashicorp_go_secure_stdlib_strutil/vendor/github.com/hashicorp/go-secure-stdlib
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/hashicorp/go-secure-stdlib)/(strutil)|\1\3/\2|1}

#-hashicorp:go-secure-stdlib:          v0.1.8:hashicorp_go_secure_stdlib_parseutil/vendor/github.com/hashicorp/go-secure-stdlib/parseutil
#+hashicorp:go-secure-stdlib:parseutil/v0.1.8:hashicorp_go_secure_stdlib_parseutil
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/hashicorp/go-secure-stdlib/(parseutil)|\1\3/\2|1}

#-hashicorp:vault:    v1.14.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault/api
#+hashicorp:vault:api/v1.14.0:hashicorp_vault_api/vendor/github.com/hashicorp/vault
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/hashicorp/vault)/(.*)|\1\3/\2|}

#-hashicorp:go-getter:v2.2.2:hashicorp_go_getter_v2/vendor/github.com/hashicorp/go-getter/gcs/v2
#+hashicorp:go-getter:v2.2.2:hashicorp_go_getter_gcs
#-hashicorp:go-getter:v2.2.2:hashicorp_go_getter_v2_1/vendor/github.com/hashicorp/go-getter/s3/v2
#+hashicorp:go-getter:v2.2.2:hashicorp_go_getter_s3
#-hashicorp:go-getter:v2.2.2:hashicorp_go_getter_v2_2/vendor/github.com/hashicorp/go-getter/v2
#+hashicorp:go-getter:v2.2.2:hashicorp_go_getter_v2
GH_TUPLE:=	${GH_TUPLE:C|:hashicorp_go_getter_.*/go-getter/([^/]+).*|:hashicorp_go_getter_\1|}

#-googleapis:gax-go:v2.12.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.12.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-ugorji:go:v1.2.6:ugorji_go_codec/vendor/github.com/ugorji/go/codec
#+ugorji:go:v1.2.6:ugorji_go_codec/vendor/github.com/ugorji/go
GH_TUPLE:=	${GH_TUPLE:S|/go/codec|/go|1}

#-google:go-genproto:49dd2c1f3d0b:google_go_genproto/vendor/google.golang.org/genproto
#+google:go-genproto:49dd2c1f3d0b:google_go_genproto
# * set - GH_TUPLE:=	${GH_TUPLE:S|/vendor/google.golang.org/genproto$||1}
#-google:go-genproto:d783a09b4405:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:d783a09b4405:google_go_genproto_rpc
# * set - GH_TUPLE:=	${GH_TUPLE:C|(:google_go_genproto_).*/|\1|}
GH_TUPLE:=	${GH_TUPLE:C|(:google_go_genproto_).*/(rpc)|\1\2|}
#-google:go-genproto:49dd2c1f3d0b:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:49dd2c1f3d0b:google_go_genproto_api
# * set - #
GH_TUPLE:=	${GH_TUPLE:N*google_go_genproto_[0-9]*}

.sinclude "${.CURDIR}/generated.deps.mk"
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

.include <bsd.port.mk>
