// automatically generated by stateify.

package verity

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (fstype *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/verity.FilesystemType"
}

func (fstype *FilesystemType) StateFields() []string {
	return []string{}
}

func (fstype *FilesystemType) beforeSave() {}

// +checklocksignore
func (fstype *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fstype.beforeSave()
}

func (fstype *FilesystemType) afterLoad() {}

// +checklocksignore
func (fstype *FilesystemType) StateLoad(stateSourceObject state.Source) {
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/verity.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"vfsfs",
		"creds",
		"allowRuntimeEnable",
		"lowerMount",
		"rootDentry",
		"alg",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.vfsfs)
	stateSinkObject.Save(1, &fs.creds)
	stateSinkObject.Save(2, &fs.allowRuntimeEnable)
	stateSinkObject.Save(3, &fs.lowerMount)
	stateSinkObject.Save(4, &fs.rootDentry)
	stateSinkObject.Save(5, &fs.alg)
}

func (fs *filesystem) afterLoad() {}

// +checklocksignore
func (fs *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.vfsfs)
	stateSourceObject.Load(1, &fs.creds)
	stateSourceObject.Load(2, &fs.allowRuntimeEnable)
	stateSourceObject.Load(3, &fs.lowerMount)
	stateSourceObject.Load(4, &fs.rootDentry)
	stateSourceObject.Load(5, &fs.alg)
}

func (i *InternalFilesystemOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/verity.InternalFilesystemOptions"
}

func (i *InternalFilesystemOptions) StateFields() []string {
	return []string{
		"RootMerkleFileName",
		"LowerName",
		"Alg",
		"RootHash",
		"AllowRuntimeEnable",
		"LowerGetFSOptions",
		"Action",
	}
}

func (i *InternalFilesystemOptions) beforeSave() {}

// +checklocksignore
func (i *InternalFilesystemOptions) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.RootMerkleFileName)
	stateSinkObject.Save(1, &i.LowerName)
	stateSinkObject.Save(2, &i.Alg)
	stateSinkObject.Save(3, &i.RootHash)
	stateSinkObject.Save(4, &i.AllowRuntimeEnable)
	stateSinkObject.Save(5, &i.LowerGetFSOptions)
	stateSinkObject.Save(6, &i.Action)
}

func (i *InternalFilesystemOptions) afterLoad() {}

// +checklocksignore
func (i *InternalFilesystemOptions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.RootMerkleFileName)
	stateSourceObject.Load(1, &i.LowerName)
	stateSourceObject.Load(2, &i.Alg)
	stateSourceObject.Load(3, &i.RootHash)
	stateSourceObject.Load(4, &i.AllowRuntimeEnable)
	stateSourceObject.Load(5, &i.LowerGetFSOptions)
	stateSourceObject.Load(6, &i.Action)
}

func (d *dentry) StateTypeName() string {
	return "pkg/sentry/fsimpl/verity.dentry"
}

func (d *dentry) StateFields() []string {
	return []string{
		"vfsd",
		"refs",
		"fs",
		"mode",
		"uid",
		"gid",
		"size",
		"parent",
		"name",
		"children",
		"childrenNames",
		"lowerVD",
		"lowerMerkleVD",
		"symlinkTarget",
		"hash",
	}
}

func (d *dentry) beforeSave() {}

// +checklocksignore
func (d *dentry) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.vfsd)
	stateSinkObject.Save(1, &d.refs)
	stateSinkObject.Save(2, &d.fs)
	stateSinkObject.Save(3, &d.mode)
	stateSinkObject.Save(4, &d.uid)
	stateSinkObject.Save(5, &d.gid)
	stateSinkObject.Save(6, &d.size)
	stateSinkObject.Save(7, &d.parent)
	stateSinkObject.Save(8, &d.name)
	stateSinkObject.Save(9, &d.children)
	stateSinkObject.Save(10, &d.childrenNames)
	stateSinkObject.Save(11, &d.lowerVD)
	stateSinkObject.Save(12, &d.lowerMerkleVD)
	stateSinkObject.Save(13, &d.symlinkTarget)
	stateSinkObject.Save(14, &d.hash)
}

// +checklocksignore
func (d *dentry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.vfsd)
	stateSourceObject.Load(1, &d.refs)
	stateSourceObject.Load(2, &d.fs)
	stateSourceObject.Load(3, &d.mode)
	stateSourceObject.Load(4, &d.uid)
	stateSourceObject.Load(5, &d.gid)
	stateSourceObject.Load(6, &d.size)
	stateSourceObject.Load(7, &d.parent)
	stateSourceObject.Load(8, &d.name)
	stateSourceObject.Load(9, &d.children)
	stateSourceObject.Load(10, &d.childrenNames)
	stateSourceObject.Load(11, &d.lowerVD)
	stateSourceObject.Load(12, &d.lowerMerkleVD)
	stateSourceObject.Load(13, &d.symlinkTarget)
	stateSourceObject.Load(14, &d.hash)
	stateSourceObject.AfterLoad(d.afterLoad)
}

func (fd *fileDescription) StateTypeName() string {
	return "pkg/sentry/fsimpl/verity.fileDescription"
}

func (fd *fileDescription) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"d",
		"isDir",
		"lowerFD",
		"lowerMappable",
		"merkleReader",
		"merkleWriter",
		"parentMerkleWriter",
		"off",
	}
}

func (fd *fileDescription) beforeSave() {}

// +checklocksignore
func (fd *fileDescription) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.d)
	stateSinkObject.Save(3, &fd.isDir)
	stateSinkObject.Save(4, &fd.lowerFD)
	stateSinkObject.Save(5, &fd.lowerMappable)
	stateSinkObject.Save(6, &fd.merkleReader)
	stateSinkObject.Save(7, &fd.merkleWriter)
	stateSinkObject.Save(8, &fd.parentMerkleWriter)
	stateSinkObject.Save(9, &fd.off)
}

func (fd *fileDescription) afterLoad() {}

// +checklocksignore
func (fd *fileDescription) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.d)
	stateSourceObject.Load(3, &fd.isDir)
	stateSourceObject.Load(4, &fd.lowerFD)
	stateSourceObject.Load(5, &fd.lowerMappable)
	stateSourceObject.Load(6, &fd.merkleReader)
	stateSourceObject.Load(7, &fd.merkleWriter)
	stateSourceObject.Load(8, &fd.parentMerkleWriter)
	stateSourceObject.Load(9, &fd.off)
}

func init() {
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*InternalFilesystemOptions)(nil))
	state.Register((*dentry)(nil))
	state.Register((*fileDescription)(nil))
}
