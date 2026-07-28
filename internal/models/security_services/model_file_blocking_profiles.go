package models

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/paloaltonetworks/terraform-provider-scm/internal/utils"
)

// Package: security_services
// This file contains models for the security_services SDK package

// FileBlockingProfiles represents the Terraform model for FileBlockingProfiles
type FileBlockingProfiles struct {
	Tfid        types.String          `tfsdk:"tfid"`
	Description basetypes.StringValue `tfsdk:"description"`
	Device      basetypes.StringValue `tfsdk:"device"`
	Folder      basetypes.StringValue `tfsdk:"folder"`
	Id          basetypes.StringValue `tfsdk:"id"`
	Name        basetypes.StringValue `tfsdk:"name"`
	Rules       basetypes.ListValue   `tfsdk:"rules"`
	Snippet     basetypes.StringValue `tfsdk:"snippet"`
}

// FileBlockingProfilesRulesInner represents a nested structure within the FileBlockingProfiles model
type FileBlockingProfilesRulesInner struct {
	Action      basetypes.StringValue `tfsdk:"action"`
	Application basetypes.ListValue   `tfsdk:"application"`
	Direction   basetypes.StringValue `tfsdk:"direction"`
	FileType    basetypes.ListValue   `tfsdk:"file_type"`
	Name        basetypes.StringValue `tfsdk:"name"`
}

// AttrTypes defines the attribute types for the FileBlockingProfiles model.
func (o FileBlockingProfiles) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"tfid":        basetypes.StringType{},
		"description": basetypes.StringType{},
		"device":      basetypes.StringType{},
		"folder":      basetypes.StringType{},
		"id":          basetypes.StringType{},
		"name":        basetypes.StringType{},
		"rules": basetypes.ListType{ElemType: basetypes.ObjectType{
			AttrTypes: map[string]attr.Type{
				"action":      basetypes.StringType{},
				"application": basetypes.ListType{ElemType: basetypes.StringType{}},
				"direction":   basetypes.StringType{},
				"file_type":   basetypes.ListType{ElemType: basetypes.StringType{}},
				"name":        basetypes.StringType{},
			},
		}},
		"snippet": basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of FileBlockingProfiles objects.
func (o FileBlockingProfiles) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// AttrTypes defines the attribute types for the FileBlockingProfilesRulesInner model.
func (o FileBlockingProfilesRulesInner) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"action":      basetypes.StringType{},
		"application": basetypes.ListType{ElemType: basetypes.StringType{}},
		"direction":   basetypes.StringType{},
		"file_type":   basetypes.ListType{ElemType: basetypes.StringType{}},
		"name":        basetypes.StringType{},
	}
}

// AttrType returns the attribute type for a list of FileBlockingProfilesRulesInner objects.
func (o FileBlockingProfilesRulesInner) AttrType() attr.Type {
	return basetypes.ObjectType{
		AttrTypes: o.AttrTypes(),
	}
}

// FileBlockingProfilesResourceSchema defines the schema for FileBlockingProfiles resource
var FileBlockingProfilesResourceSchema = schema.Schema{
	MarkdownDescription: "FileBlockingProfile resource",
	Attributes: map[string]schema.Attribute{
		"description": schema.StringAttribute{
			MarkdownDescription: "Description",
			Optional:            true,
		},
		"device": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("folder"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"folder": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("snippet"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
				utils.FolderValidator(),
			},
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"id": schema.StringAttribute{
			MarkdownDescription: "The UUID of the file blocking profile",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the file blocking profile",
			Required:            true,
		},
		"rules": schema.ListNestedAttribute{
			MarkdownDescription: "A list of file blocking rules",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("alert", "block", "continue"),
						},
						MarkdownDescription: "The action to take when the rule match criteria is met. Possible values are `alert`, `block` and `continue`.",
						Required:            true,
					},
					"application": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "The application transferring the files (App-ID naming). Possible values are `any`, `axifile`, `dl-free`, `facebook-mail`, `file.io`, `giphy-base`, `glassdoor-uploading`, `http-proxy`, `redbooth`, `send-anywhere` and `zoho-mail`.",
						Required:            true,
					},
					"direction": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("download", "upload", "both"),
						},
						MarkdownDescription: "The direction of the file transfer. Possible values are `download`, `upload` and `both`.",
						Required:            true,
					},
					"file_type": schema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "The file type. Possible values are `7z`, `Multi-Level-Encoding`, `PE`, `access-shortcut`, `ace`, `ade`, `adp`, `ai`, `aip-encrypted-docx`, `aip-encrypted-pptx`, `aip-encrypted-xlsx`, `any`, `apk`, `arj`, `asp`, `avi`, `avi-divx`, `avi-xvid`, `bas`, `bat`, `bmp`, `bmp-upload`, `bzip2`, `cab`, `catpart`, `cdr`, `chm`, `cin`, `class`, `cmd`, `com`, `cpl`, `csv`, `deflate64-zip`, `der`, `dll`, `dmg`, `doc`, `docm`, `docx`, `dpx`, `dsn`, `dwf`, `dwg`, `dxf`, `edif`, `elf`, `emf`, `encrypted-7z`, `encrypted-doc`, `encrypted-docx`, `encrypted-office2007`, `encrypted-pdf`, `encrypted-ppt`, `encrypted-pptx`, `encrypted-rar`, `encrypted-xls`, `encrypted-xlsx`, `encrypted-zip`, `eps`, `exe`, `exr`, `flash`, `flv`, `gds`, `gif`, `gif-upload`, `gzip`, `hlp`, `hta`, `hwp`, `hwpx`, `ichitaro`, `iff`, `inf`, `ins`, `iqy`, `iso`, `its`, `iwork-keynote`, `iwork-numbers`, `iwork-pages`, `jar`, `jpeg`, `jpeg-upload`, `js`, `jse`, `lnk`, `lzh`, `ma`, `macapp`, `mach-o`, `mb`, `mda`, `mdb`, `mdi`, `mdt`, `mdw`, `mdz`, `mht`, `microsoft-shell`, `mif`, `mkv`, `mov`, `mp3`, `mp4`, `mpeg`, `mpeg-ts`, `mpkg`, `msc`, `msi`, `msoffice`, `msp`, `ocx`, `pbix`, `pbm`, `pcl`, `pdf`, `pem`, `pgp`, `pif`, `pkg`, `pl`, `png`, `png-upload`, `powershell`, `ppt`, `pptx`, `prg`, `psd`, `py`, `rar`, `reg`, `renamed-zip`, `rla`, `rm`, `rpf`, `rtf`, `scf`, `scr`, `sgi`, `sh`, `shk`, `shs`, `slk`, `softimg`, `split-cab`, `split-rar`, `stp`, `svg`, `sys`, `tar`, `tdb`, `tif`, `tiff`, `tmp`, `torrent`, `url`, `vb`, `vbe`, `vbs`, `vxd`, `webm`, `wmf`, `wmv`, `wri`, `wsf`, `wsh`, `xll`, `xls`, `xlsx`, `xpm`, `zcompressed` and `zip`.",
						Required:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: "The name of the file blocking rule",
						Required:            true,
					},
				},
			},
		},
		"snippet": schema.StringAttribute{
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("device"),
					path.MatchRelative().AtParent().AtName("folder"),
				),
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z\\d\\-_\\. ]+$"), "pattern must match "+"^[a-zA-Z\\d\\-_\\. ]+$"),
			},
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"tfid": schema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
	},
}

// FileBlockingProfilesDataSourceSchema defines the schema for FileBlockingProfiles data source
var FileBlockingProfilesDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "FileBlockingProfile data source",
	Attributes: map[string]dsschema.Attribute{
		"description": dsschema.StringAttribute{
			MarkdownDescription: "Description",
			Computed:            true,
		},
		"device": dsschema.StringAttribute{
			MarkdownDescription: "The device in which the resource is defined",
			Optional:            true,
			Computed:            true,
		},
		"folder": dsschema.StringAttribute{
			MarkdownDescription: "The folder in which the resource is defined",
			Optional:            true,
			Computed:            true,
		},
		"id": dsschema.StringAttribute{
			MarkdownDescription: "The UUID of the file blocking profile",
			Required:            true,
		},
		"name": dsschema.StringAttribute{
			MarkdownDescription: "The name of the file blocking profile",
			Optional:            true,
			Computed:            true,
		},
		"rules": dsschema.ListNestedAttribute{
			MarkdownDescription: "A list of file blocking rules",
			Computed:            true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: map[string]dsschema.Attribute{
					"action": dsschema.StringAttribute{
						MarkdownDescription: "The action to take when the rule match criteria is met. Possible values are `alert`, `block` and `continue`.",
						Computed:            true,
					},
					"application": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "The application transferring the files (App-ID naming). Possible values are `any`, `axifile`, `dl-free`, `facebook-mail`, `file.io`, `giphy-base`, `glassdoor-uploading`, `http-proxy`, `redbooth`, `send-anywhere` and `zoho-mail`.",
						Computed:            true,
					},
					"direction": dsschema.StringAttribute{
						MarkdownDescription: "The direction of the file transfer. Possible values are `download`, `upload` and `both`.",
						Computed:            true,
					},
					"file_type": dsschema.ListAttribute{
						ElementType:         types.StringType,
						MarkdownDescription: "The file type. Possible values are `7z`, `Multi-Level-Encoding`, `PE`, `access-shortcut`, `ace`, `ade`, `adp`, `ai`, `aip-encrypted-docx`, `aip-encrypted-pptx`, `aip-encrypted-xlsx`, `any`, `apk`, `arj`, `asp`, `avi`, `avi-divx`, `avi-xvid`, `bas`, `bat`, `bmp`, `bmp-upload`, `bzip2`, `cab`, `catpart`, `cdr`, `chm`, `cin`, `class`, `cmd`, `com`, `cpl`, `csv`, `deflate64-zip`, `der`, `dll`, `dmg`, `doc`, `docm`, `docx`, `dpx`, `dsn`, `dwf`, `dwg`, `dxf`, `edif`, `elf`, `emf`, `encrypted-7z`, `encrypted-doc`, `encrypted-docx`, `encrypted-office2007`, `encrypted-pdf`, `encrypted-ppt`, `encrypted-pptx`, `encrypted-rar`, `encrypted-xls`, `encrypted-xlsx`, `encrypted-zip`, `eps`, `exe`, `exr`, `flash`, `flv`, `gds`, `gif`, `gif-upload`, `gzip`, `hlp`, `hta`, `hwp`, `hwpx`, `ichitaro`, `iff`, `inf`, `ins`, `iqy`, `iso`, `its`, `iwork-keynote`, `iwork-numbers`, `iwork-pages`, `jar`, `jpeg`, `jpeg-upload`, `js`, `jse`, `lnk`, `lzh`, `ma`, `macapp`, `mach-o`, `mb`, `mda`, `mdb`, `mdi`, `mdt`, `mdw`, `mdz`, `mht`, `microsoft-shell`, `mif`, `mkv`, `mov`, `mp3`, `mp4`, `mpeg`, `mpeg-ts`, `mpkg`, `msc`, `msi`, `msoffice`, `msp`, `ocx`, `pbix`, `pbm`, `pcl`, `pdf`, `pem`, `pgp`, `pif`, `pkg`, `pl`, `png`, `png-upload`, `powershell`, `ppt`, `pptx`, `prg`, `psd`, `py`, `rar`, `reg`, `renamed-zip`, `rla`, `rm`, `rpf`, `rtf`, `scf`, `scr`, `sgi`, `sh`, `shk`, `shs`, `slk`, `softimg`, `split-cab`, `split-rar`, `stp`, `svg`, `sys`, `tar`, `tdb`, `tif`, `tiff`, `tmp`, `torrent`, `url`, `vb`, `vbe`, `vbs`, `vxd`, `webm`, `wmf`, `wmv`, `wri`, `wsf`, `wsh`, `xll`, `xls`, `xlsx`, `xpm`, `zcompressed` and `zip`.",
						Computed:            true,
					},
					"name": dsschema.StringAttribute{
						MarkdownDescription: "The name of the file blocking rule",
						Computed:            true,
					},
				},
			},
		},
		"snippet": dsschema.StringAttribute{
			MarkdownDescription: "The snippet in which the resource is defined",
			Optional:            true,
			Computed:            true,
		},
		"tfid": dsschema.StringAttribute{
			MarkdownDescription: "The Terraform ID.",
			Computed:            true,
		},
	},
}

// FileBlockingProfilesListModel represents the data model for a list data source.
type FileBlockingProfilesListModel struct {
	Tfid    types.String           `tfsdk:"tfid"`
	Data    []FileBlockingProfiles `tfsdk:"data"`
	Limit   types.Int64            `tfsdk:"limit"`
	Offset  types.Int64            `tfsdk:"offset"`
	Name    types.String           `tfsdk:"name"`
	Total   types.Int64            `tfsdk:"total"`
	Folder  types.String           `tfsdk:"folder"`
	Snippet types.String           `tfsdk:"snippet"`
	Device  types.String           `tfsdk:"device"`
}

// FileBlockingProfilesListDataSourceSchema defines the schema for a list data source.
var FileBlockingProfilesListDataSourceSchema = dsschema.Schema{
	MarkdownDescription: "Retrieves a listing of config items.",
	Attributes: map[string]dsschema.Attribute{
		"tfid": dsschema.StringAttribute{Description: "The Terraform ID.", Computed: true},
		"data": dsschema.ListNestedAttribute{
			Description: "The data.",
			Computed:    true,
			NestedObject: dsschema.NestedAttributeObject{
				Attributes: FileBlockingProfilesDataSourceSchema.Attributes,
			},
		},
		"limit":   dsschema.Int64Attribute{Description: "The max number of items to return. Default: 200.", Optional: true},
		"offset":  dsschema.Int64Attribute{Description: "The offset of the first item to return.", Optional: true},
		"name":    dsschema.StringAttribute{Description: "The name of the item.", Optional: true},
		"total":   dsschema.Int64Attribute{Description: "The total number of items.", Computed: true},
		"folder":  dsschema.StringAttribute{Description: "The folder of the item. Default: Shared.", Optional: true},
		"snippet": dsschema.StringAttribute{Description: "The snippet of the item.", Optional: true},
		"device":  dsschema.StringAttribute{Description: "The device of the item.", Optional: true},
	},
}
