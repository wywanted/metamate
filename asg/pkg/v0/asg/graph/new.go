package graph

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/enumnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/cardinality"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/preposition"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/verbs/infinitive"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/verbs/past"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/verbs/present"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
)

const (
	String  = "string"
	Int32   = "int32"
	Float64 = "float64"
	Bool    = "bool"
)

func NewRoot() (root *RootNode) {
	root = NewRootNode()

	addEntities(root)

	root.Types.Each(func(tn *TypeNode) {
		tn.Flags().Set(typeflags.IsEntity, true)
	})

	addExpansion(root)

	root.Types.Flagged(typeflags.IsEntity, false).Each(func(tn *TypeNode) {
		tn.Flags().Set(typeflags.IsExpansion, true)
	})

	addActionEndpoints(root)

	return
}

func addActionEndpoints(root *RootNode) {
	root.AddActionEndpoint(endpointnames.LookupService,
		FieldNodeSlice{},
		FieldNodeSlice{
			TypeField("service", typenames.Service),
		},
	)
}

func addExpansion(root *RootNode) {
	root.AddTypeNode(typenames.IdGetMode, FieldNodeSlice{
		TypeField(fieldnames.Id, typenames.Id),
	})

	root.AddTypeNode(typenames.RelationGetMode, FieldNodeSlice{
		TypeField(fieldnames.Id, typenames.ServiceId),
		StringField(fieldnames.Relation),
	})

	root.AddTypeNode(typenames.SearchGetMode, FieldNodeSlice{
		StringField("term"),
		TypeField("location", typenames.LocationQuery),
	})

	root.AddTypeNode(typenames.CollectionGetMode, FieldNodeSlice{
		ListField(TypeField("pages", typenames.ServicePage)),
	})

	root.AddUnion(typenames.GetMode, []interface{}{
		TypeField(fieldnames.Collection, typenames.CollectionGetMode),
		TypeField(fieldnames.Id, typenames.Id),
		TypeField(fieldnames.Relation, typenames.RelationGetMode),
		TypeField(fieldnames.Search, typenames.SearchGetMode),
	})

	root.AddUnion(typenames.PipeMode, []interface{}{
		TypeField(fieldnames.Context, typenames.ContextPipeMode),
	})

	root.AddTypeNode(typenames.ContextPipeMode, FieldNodeSlice{
		EnumField(fieldnames.Stage, enumnames.RequestStage),
		EnumField(fieldnames.Method, enumnames.Methods),
		EnumField(fieldnames.Requester, enumnames.BusActor),
	})

	root.AddEnumNode(enumnames.RequestStage, []string{
		"request",
		"response",
	})

	root.AddEnumNode(enumnames.BusActor, []string{
		"client",
		"bus",
		"service",
	})

	root.AddEnumNode(enumnames.Methods, []string{
		"get",
		"pipe",
		"action",
	})

	root.AddEnumNode(enumnames.ErrorKind, []string{
		"service",
		"upstream",
		"pipe",
		"internal",
		"requestValidation",
		"responseValidation",
		"idNotPresent",
		"idAlreadyPresent",
		"noServiceMatch",
	})

	root.AddTypeNode(typenames.Error, FieldNodeSlice{
		EnumField(fieldnames.Kind, enumnames.ErrorKind),
		StringField(fieldnames.Message),
		TypeField(fieldnames.Id, typenames.Id),
		TypeField(fieldnames.Service, typenames.Service),
		TypeField(fieldnames.Wraps, typenames.Error),
	})

	root.AddTypeNode(typenames.Warning, FieldNodeSlice{
		StringField(fieldnames.Message),
	})

	root.AddTypeNode(typenames.RequestMeta, FieldNodeSlice{
		TypeField("createdAt", typenames.Timestamp),
	})

	root.AddTypeNode(typenames.IndexPage, FieldNodeSlice{
		Int32Field(fieldnames.Value),
	})

	root.AddTypeNode(typenames.CursorPage, FieldNodeSlice{
		StringField(fieldnames.Value),
	})

	root.AddTypeNode(typenames.OffsetPage, FieldNodeSlice{
		Int32Field("offset"),
		Int32Field("limit"),
	})

	root.AddUnion(typenames.Page, []interface{}{
		typenames.IndexPage,
		typenames.OffsetPage,
		typenames.CursorPage,
	})

	root.AddTypeNode(typenames.ServicePage, FieldNodeSlice{
		TypeField(fieldnames.Id, typenames.ServiceId),
		TypeField("page", typenames.Page),
	})

	root.AddTypeNode(typenames.Pagination, FieldNodeSlice{
		ListField(TypeField("previous", typenames.ServicePage)),
		ListField(TypeField("current", typenames.ServicePage)),
		ListField(TypeField("next", typenames.ServicePage)),
	})

	root.AddEnumNode(enumnames.SortKind, []string{
		"asc",
		"desc",
	})

	return
}

func addEntities(root *RootNode) {
	root.AddTypeNode(typenames.TypeMeta, FieldNodeSlice{
		TypeField(fieldnames.Service, typenames.Service),
		BoolField("archived"),
		BoolField("sensitive"),
		TypeField("createdAt", typenames.Timestamp),
		TypeField("updateAt", typenames.Timestamp),
		TypeField("deletedAt", typenames.Timestamp),
	})

	root.AddEnumNode(enumnames.LengthUnit, []string{
		"nm",
		"mcm",
		"mm",
		"cm",
		"dm",
		"m",
		"km",
		"th",
		"in",
		"ft",
		"yd",
		"mi",
		"lea",
	})

	root.AddValue(typenames.LengthValue, enumnames.LengthUnit)

	root.AddScalar(typenames.LengthScalar, enumnames.LengthUnit)

	root.AddTypeNode(typenames.LocationQuery, FieldNodeSlice{
		TypeField("radiusLt", typenames.LengthScalar),
		StringField("street"),
		StringField("zipCode"),
		StringField("city"),
		StringField("cityDistrict"),
		StringField("country"),
		StringField("countryState"),
		StringField("countryStateDistrict"),
	})

	root.AddEnumNode(enumnames.DurationUnit, []string{
		"ns",
		"ms",
		"s",
		"m",
		"h",
		"d",
		"w",
		"y",
	})

	root.AddScalar(typenames.DurationScalar, enumnames.DurationUnit)

	root.AddEnumNode(enumnames.Language, []string{
		"en",
	})

	root.AddEnumNode(enumnames.FormattingKind, []string{
		"html",
		"plain",
		"markdown",
	})

	root.AddTypeNode(typenames.Text, FieldNodeSlice{
		EnumField("language", enumnames.Language),
		EnumField("formatting", enumnames.FormattingKind),
		StringField(fieldnames.Value),
		//ListField(TypeField("translations", Translation)),
	})

	root.AddEnumNode(enumnames.ValueKind, []string{
		fieldnames.Value,
		"range",
	})

	root.AddTypeNode(typenames.FloatRange, FieldNodeSlice{
		Float64Field("from"),
		Float64Field("to"),
	})

	root.AddTypeNode(typenames.Service, FieldNodeSlice{
		BoolField(fieldnames.IsVirtual),
		StringField(fieldnames.Name),
		StringField(fieldnames.SdkVersion),
		TypeField(fieldnames.Url, typenames.Url),
		Int32Field(fieldnames.Port),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.Url, FieldNodeSlice{
		StringField(fieldnames.Value, Flags{fieldflags.ValidateUrl: true}),
	}, Flags{
		typeflags.GetPassEndpoint: true,
	})

	root.AddTypeNode(typenames.Email, FieldNodeSlice{
		StringField(fieldnames.Value, Flags{fieldflags.ValidateEmail: true, fieldflags.ValidateIsSet: true}),
	})

	root.AddTypeNode(typenames.ServiceId, FieldNodeSlice{
		StringField(fieldnames.Value, Flags{fieldflags.ValidateIsSet: true}),
		StringField(fieldnames.ServiceName),
	})

	root.AddUnion(typenames.Id, []interface{}{
		typenames.Url,
		StringField(fieldnames.Name),
		StringField("username"),
		StringField("ean"),
		TypeField(fieldnames.ServiceId, typenames.ServiceId),
		StringField(fieldnames.Local),
		BoolField(fieldnames.Me),
		typenames.Email,
	})

	root.AddUnion(typenames.Timestamp, []interface{}{
		TypeField("unix", typenames.DurationScalar),
	})

	root.AddTypeNode(typenames.Image, FieldNodeSlice{
		TypeField("url", typenames.Url),
		BoolField("isPreview"),
		Int32Field("width"),
		Int32Field("height"),
		TypeField("description", typenames.Text),
	})

	root.AddTypeNode(typenames.HyperLink, FieldNodeSlice{
		TypeField("url", typenames.Url),
		StringField("label"),
	})

	root.AddTypeNode(typenames.SocialAccount, FieldNodeSlice{
		StringField("username"),
		StringField("displayName"),
		TypeField("note", typenames.Text),
		Int32Field("points"),
		TypeField("avatar", typenames.Image),
		TypeField("header", typenames.Image),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.Attachment, FieldNodeSlice{
		StringField("description"),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.Info, FieldNodeSlice{
		TypeField("name", typenames.Text),
		TypeField("description", typenames.Text),
		TypeField("purpose", typenames.Text),
	})

	root.AddEnumNode(enumnames.PostFeedKind, []string{
		"channel",
		"privateChannel",
		"conversation",
	})

	root.AddTypeNode(typenames.PostFeed, FieldNodeSlice{
		TypeField("info", typenames.Info),
		EnumField("kind", enumnames.PostFeedKind),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddEnumNode(enumnames.PostKind, []string{
		"post",
		"reply",
	})

	root.AddTypeNode(typenames.Post, FieldNodeSlice{
		EnumField("kind", enumnames.PostKind),
		TypeField("title", typenames.Text),
		TypeField("content", typenames.Text),
		TypeField("spoilerText", typenames.Text),
		Int32Field("totalWasRepliedToByPostsCount"),
		BoolField("isSensitive"),
		BoolField("isPinned"),
		ListField(TypeField("links", typenames.HyperLink)),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Authors}, cardinality.Many, typenames.Post},
		RelationPath{typenames.Post, []string{past.Authored, preposition.By}, cardinality.One, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Follows}, cardinality.Many, typenames.SocialAccount},
		RelationPath{typenames.SocialAccount, []string{past.Followed, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Mutes}, cardinality.Many, typenames.SocialAccount},
		RelationPath{typenames.SocialAccount, []string{past.Muted, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Requests, preposition.To, infinitive.Follow}, cardinality.Many, typenames.SocialAccount},
		RelationPath{typenames.SocialAccount, []string{past.Requested, preposition.To, infinitive.Be, past.Followed, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Blocks}, cardinality.Many, typenames.SocialAccount},
		RelationPath{typenames.SocialAccount, []string{past.Blocked, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.Post, []string{present.Replies, preposition.To}, cardinality.One, typenames.SocialAccount},
		RelationPath{typenames.SocialAccount, []string{past.Was, past.Replied, preposition.To, preposition.By}, cardinality.Many, typenames.Post},
	)

	root.AddRelationNode(
		RelationPath{typenames.Post, []string{present.Replies, preposition.To}, cardinality.One, typenames.Post},
		RelationPath{typenames.Post, []string{past.Was, past.Replied, preposition.To, preposition.By}, cardinality.Many, typenames.Post},
	)

	root.AddRelationNode(
		RelationPath{typenames.Post, []string{present.Reblogs}, cardinality.One, typenames.Post},
		RelationPath{typenames.Post, []string{past.Reblogged, preposition.By}, cardinality.Many, typenames.Post},
	)

	root.AddRelationNode(
		RelationPath{typenames.Post, []string{present.Mentions}, cardinality.Many, typenames.SocialAccount},
		RelationPath{typenames.SocialAccount, []string{past.Mentioned, preposition.By}, cardinality.Many, typenames.Post},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Favors}, cardinality.Many, typenames.Post},
		RelationPath{typenames.Post, []string{past.Favored, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.Post, []string{present.Attaches}, cardinality.Many, typenames.Attachment},
		RelationPath{typenames.Attachment, []string{past.Attached, preposition.To}, cardinality.One, typenames.Post},
	)

	root.AddRelationNode(
		RelationPath{typenames.PostFeed, []string{present.Contains}, cardinality.Many, typenames.Post},
		RelationPath{typenames.Post, []string{past.Contained, preposition.By}, cardinality.Many, typenames.PostFeed},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Participates}, cardinality.Many, typenames.PostFeed},
		RelationPath{typenames.PostFeed, []string{past.Participated, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{past.Read}, cardinality.Many, typenames.Post},
		RelationPath{typenames.Post, []string{past.Read, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{preposition.Not, past.Read}, cardinality.Many, typenames.Post},
		RelationPath{typenames.Post, []string{preposition.Not, past.Read, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Reblogs}, cardinality.Many, typenames.Post},
		RelationPath{typenames.Post, []string{past.Reblogged, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Mutes}, cardinality.Many, typenames.Post},
		RelationPath{typenames.Post, []string{past.Muted, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.SocialAccount, []string{present.Bookmarks}, cardinality.Many, typenames.Post},
		RelationPath{typenames.Post, []string{past.Bookmarked, preposition.By}, cardinality.Many, typenames.SocialAccount},
	)

	root.AddRelationNode(
		RelationPath{typenames.Dummy, []string{present.Knows}, cardinality.Many, typenames.Dummy},
		RelationPath{typenames.Dummy, []string{past.Knew, preposition.By}, cardinality.Many, typenames.Dummy},
	)

	root.AddRelationNode(
		RelationPath{typenames.Dummy, []string{present.Knows}, cardinality.Many, typenames.BlueDummy},
		RelationPath{typenames.BlueDummy, []string{past.Knew, preposition.By}, cardinality.Many, typenames.Dummy},
	)

	root.AddEnumNode(enumnames.DummyKind, []string{
		"red",
		"blue",
		"green",
	})

	root.AddTypeNode(typenames.Dummy, FieldNodeSlice{
		EnumField("enumField", enumnames.DummyKind),
		StringField("stringField"),
		Int32Field("int32Field"),
		Float64Field("float64Field"),
		BoolField("boolField"),
		TypeField("unionField", typenames.DummyUnion),
		ListField(EnumField("enumList", enumnames.DummyKind)),
		ListField(StringField("stringList")),
		ListField(Int32Field("int32List")),
		ListField(Float64Field("float64List")),
		ListField(BoolField("boolList")),
		//ListField(TypeField("typeList", Dummy)),
		ListField(TypeField("unionList", typenames.DummyUnion)),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddTypeNode(typenames.BlueDummy, FieldNodeSlice{
		EnumField("enumField", enumnames.DummyKind),
		StringField("stringField"),
		Int32Field("int32Field"),
		Float64Field("float64Field"),
		BoolField("boolField"),
		TypeField("unionField", typenames.DummyUnion),
		ListField(EnumField("enumList", enumnames.DummyKind)),
		ListField(StringField("stringList")),
		ListField(Int32Field("int32List")),
		ListField(Float64Field("float64List")),
		ListField(BoolField("boolList")),
		//ListField(TypeField("typeList", Dummy)),
		ListField(TypeField("unionList", typenames.DummyUnion)),
	}, Flags{
		typeflags.GetEndpoints: true,
	})

	root.AddUnion(typenames.DummyUnion, []interface{}{
		StringField("stringField"),
		Int32Field("int32Field"),
		Float64Field("float64Field"),
		BoolField("boolField"),
		EnumField("enumField", enumnames.DummyKind),
	})

	root.AddEnumNode(enumnames.CurrencyUnit, []string{
		"eur",
	})

	root.AddScalar(typenames.CurrencyScalar, enumnames.CurrencyUnit)
}
