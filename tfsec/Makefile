PORTNAME=	tfsec
PORTVERSION=	1.28.14
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	nulldutra@proton.me
COMMENT=	Security scanner for Terraform
WWW=		https://aquasecurity.github.io/${PORTNAME}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aquasecurity

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/iam
	@test -z ${WRKSRC_googleapis_google_cloud_go_iam} || ${RLN} ${WRKSRC_googleapis_google_cloud_go_iam}/iam ${WRKSRC}/vendor/cloud.google.com/go
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/storage
	@test -z ${WRKSRC_googleapis_google_cloud_go_storage} || ${RLN} ${WRKSRC_googleapis_google_cloud_go_storage}/storage ${WRKSRC}/vendor/cloud.google.com/go
	@${MKDIR} ${WRKSRC}/vendor/cloud.google.com/go/compute
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/compute/metadata
	@${RLN} ${WRKSRC_googleapis_google_cloud_go_compute_metadata}/compute/metadata ${WRKSRC}/vendor/cloud.google.com/go/compute
	@${MKDIR} ${WRKSRC}/vendor/gopkg.in
	@test -z ${WRKSRC_go_ini_ini} || ${RLN} ${WRKSRC_go_ini_ini} ${WRKSRC}/vendor/gopkg.in/ini.v1
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/api
	@test -z ${WRKSRC_google_go_genproto_api} || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@${RM} -r ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/rpc
	@test -z ${WRKSRC_google_go_genproto_rpc} || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=imdario GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=google-cloud-go GHT_PATH=cloud.google.com/go/compute/metadata )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-go-ini:ini:v1.67.0:go_ini_ini_1/vendor/gopkg.in/ini.v1
GH_TUPLE:=	${GH_TUPLE:Ngo-ini\:ini\:*/ini.v1}

#-googleapis:gax-go:v2.12.3:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.12.3:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|}

# google:go-genproto:012b6fc9bca9:google_go_genproto/vendor/google.golang.org/genproto
#-google:go-genproto:f6361c86f094:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:f6361c86f094:google_go_genproto_api
#-google:go-genproto:f6361c86f094:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:f6361c86f094:google_go_genproto_rpc
GH_TUPLE:=	${GH_TUPLE:C|_[[:digit:]]/vendor/google.golang.org/genproto/googleapis/|_|}

#+googleapis:google-cloud-go:iam/v1.1.6:googleapis_google_cloud_go_iam
#+googleapis:google-cloud-go:storage/v1.38.0:googleapis_google_cloud_go_storage
# googleapis:google-cloud-go:v0.112.1:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v1.1.6:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.38.0:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/cloud.google.com/go/(iam\|storage)|\1\3/\2|}

#-owenrumney:go-sarif:v2.1.2:owenrumney_go_sarif_v2/vendor/github.com/owenrumney/go-sarif/v2
#+owenrumney:go-sarif:v2.1.2:owenrumney_go_sarif_v2/vendor/github.com/owenrumney/go-sarif
GH_TUPLE:=	${GH_TUPLE:S|/go-sarif/v2|/go-sarif|}

#-census-instrumentation:opencensus-go:v0.24.0:census_instrumentation_opencensus_go/vendor/go.opencensus.io
GH_TUPLE:=	${GH_TUPLE:Ncensus-instrumentation\:opencensus-go\:*}

.sinclude "${.CURDIR}/generated.deps.mk"
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

.include <bsd.port.mk>
