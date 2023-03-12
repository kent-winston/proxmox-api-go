package proxmox

type QemuSataDisk struct {
	AsyncIO    QemuDiskAsyncIO   `json:"asyncio,omitempty"`
	Backup     bool              `json:"backup,omitempty"`
	Bandwidth  QemuDiskBandwidth `json:"bandwith,omitempty"`
	Cache      QemuDiskCache     `json:"cache,omitempty"`
	Discard    bool              `json:"discard,omitempty"`
	EmulateSSD bool              `json:"emulatessd,omitempty"`
	Format     *QemuDiskFormat   `json:"format,omitempty"`
	Id         *uint             `json:"id,omitempty"`
	Replicate  bool              `json:"replicate,omitempty"`
	Serial     QemuDiskSerial    `json:"serial,omitempty"`
	Size       uint              `json:"size,omitempty"`
	Storage    string            `json:"storage,omitempty"`
}

// TODO write test
func (disk QemuSataDisk) mapToApiValues(vmID uint, create bool) string {
	return qemuDisk{
		AsyncIO:    disk.AsyncIO,
		Backup:     disk.Backup,
		Bandwidth:  disk.Bandwidth,
		Cache:      disk.Cache,
		Discard:    disk.Discard,
		EmulateSSD: disk.EmulateSSD,
		Replicate:  disk.Replicate,
		Serial:     disk.Serial,
		Size:       disk.Size,
		Storage:    disk.Storage,
		Type:       sata,
	}.mapToApiValues(vmID, create)
}

type QemuSataDisks struct {
	Disk_0 *QemuSataStorage `json:"0,omitempty"`
	Disk_1 *QemuSataStorage `json:"1,omitempty"`
	Disk_2 *QemuSataStorage `json:"2,omitempty"`
	Disk_3 *QemuSataStorage `json:"3,omitempty"`
	Disk_4 *QemuSataStorage `json:"4,omitempty"`
	Disk_5 *QemuSataStorage `json:"5,omitempty"`
}

// TODO write test
func (disks QemuSataDisks) mapToApiValues(currentDisks *QemuSataDisks, vmID uint, params map[string]interface{}, changes *qemuUpdateChanges) {
	tmpCurrentDisks := QemuSataDisks{}
	if currentDisks != nil {
		tmpCurrentDisks = *currentDisks
	}
	disks.Disk_0.convertDataStructure().markDiskChanges(tmpCurrentDisks.Disk_0.convertDataStructure(), vmID, "sata0", params, changes)
	disks.Disk_1.convertDataStructure().markDiskChanges(tmpCurrentDisks.Disk_1.convertDataStructure(), vmID, "sata1", params, changes)
	disks.Disk_2.convertDataStructure().markDiskChanges(tmpCurrentDisks.Disk_2.convertDataStructure(), vmID, "sata2", params, changes)
	disks.Disk_3.convertDataStructure().markDiskChanges(tmpCurrentDisks.Disk_3.convertDataStructure(), vmID, "sata3", params, changes)
	disks.Disk_4.convertDataStructure().markDiskChanges(tmpCurrentDisks.Disk_4.convertDataStructure(), vmID, "sata4", params, changes)
	disks.Disk_5.convertDataStructure().markDiskChanges(tmpCurrentDisks.Disk_5.convertDataStructure(), vmID, "sata5", params, changes)
}

// TODO write test
func (QemuSataDisks) mapToStruct(params map[string]interface{}) *QemuSataDisks {
	disks := QemuSataDisks{}
	var structPopulated bool
	if _, isSet := params["sata0"]; isSet {
		disks.Disk_0 = QemuSataStorage{}.mapToStruct(params["sata0"].(string))
		structPopulated = true
	}
	if _, isSet := params["sata1"]; isSet {
		disks.Disk_1 = QemuSataStorage{}.mapToStruct(params["sata1"].(string))
		structPopulated = true
	}
	if _, isSet := params["sata2"]; isSet {
		disks.Disk_2 = QemuSataStorage{}.mapToStruct(params["sata2"].(string))
		structPopulated = true
	}
	if _, isSet := params["sata3"]; isSet {
		disks.Disk_3 = QemuSataStorage{}.mapToStruct(params["sata3"].(string))
		structPopulated = true
	}
	if _, isSet := params["sata4"]; isSet {
		disks.Disk_4 = QemuSataStorage{}.mapToStruct(params["sata4"].(string))
		structPopulated = true
	}
	if _, isSet := params["sata5"]; isSet {
		disks.Disk_5 = QemuSataStorage{}.mapToStruct(params["sata5"].(string))
		structPopulated = true
	}
	if structPopulated {
		return &disks
	}
	return nil
}

type QemuSataPassthrough struct {
	AsyncIO    QemuDiskAsyncIO
	Backup     bool
	Bandwidth  QemuDiskBandwidth
	Cache      QemuDiskCache
	Discard    bool
	EmulateSSD bool
	File       string
	Replicate  bool
	Serial     QemuDiskSerial `json:"serial,omitempty"`
	Size       uint           //size is only returned and setting it has no effect
}

// TODO write test
func (passthrough QemuSataPassthrough) mapToApiValues() string {
	return qemuDisk{
		AsyncIO:    passthrough.AsyncIO,
		Backup:     passthrough.Backup,
		Bandwidth:  passthrough.Bandwidth,
		Cache:      passthrough.Cache,
		Discard:    passthrough.Discard,
		EmulateSSD: passthrough.EmulateSSD,
		File:       passthrough.File,
		Replicate:  passthrough.Replicate,
		Serial:     passthrough.Serial,
		Type:       sata,
	}.mapToApiValues(0, false)
}

type QemuSataStorage struct {
	CdRom       *QemuCdRom
	CloudInit   *QemuCloudInitDisk
	Disk        *QemuSataDisk
	Passthrough *QemuSataPassthrough
}

// TODO write test
// converts to qemuStorage
func (storage *QemuSataStorage) convertDataStructure() *qemuStorage {
	if storage == nil {
		return nil
	}
	generalizedStorage := qemuStorage{
		CdRom:     storage.CdRom,
		CloudInit: storage.CloudInit,
	}
	if storage.Disk != nil {
		generalizedStorage.Disk = &qemuDisk{
			AsyncIO:    storage.Disk.AsyncIO,
			Backup:     storage.Disk.Backup,
			Bandwidth:  storage.Disk.Bandwidth,
			Cache:      storage.Disk.Cache,
			Discard:    storage.Disk.Discard,
			EmulateSSD: storage.Disk.EmulateSSD,
			Format:     storage.Disk.Format,
			Id:         storage.Disk.Id,
			Replicate:  storage.Disk.Replicate,
			Serial:     storage.Disk.Serial,
			Size:       storage.Disk.Size,
			Storage:    storage.Disk.Storage,
		}
	}
	if storage.Passthrough != nil {
		generalizedStorage.Passthrough = &qemuDisk{
			AsyncIO:    storage.Passthrough.AsyncIO,
			Backup:     storage.Passthrough.Backup,
			Bandwidth:  storage.Passthrough.Bandwidth,
			Cache:      storage.Passthrough.Cache,
			Discard:    storage.Passthrough.Discard,
			EmulateSSD: storage.Passthrough.EmulateSSD,
			File:       storage.Passthrough.File,
			Replicate:  storage.Passthrough.Replicate,
			Serial:     storage.Passthrough.Serial,
		}
	}
	return &generalizedStorage
}

// TODO write test
func (storage QemuSataStorage) mapToApiValues(vmID uint, create bool) string {
	if storage.Disk != nil {
		return storage.Disk.mapToApiValues(vmID, create)
	}
	if storage.CdRom != nil {
		return storage.CdRom.mapToApiValues()
	}
	if storage.CloudInit != nil {
		return storage.CloudInit.mapToApiValues()
	}
	if storage.Passthrough != nil {
		return storage.Passthrough.mapToApiValues()
	}
	return ""
}

// TODO write test
func (QemuSataStorage) mapToStruct(param string) *QemuSataStorage {
	settings := splitStringOfSettings(param)
	tmpCdRom := qemuCdRom{}.mapToStruct(settings)
	if tmpCdRom != nil {
		if tmpCdRom.FileType == "" {
			return &QemuSataStorage{CdRom: QemuCdRom{}.mapToStruct(*tmpCdRom)}
		} else {
			return &QemuSataStorage{CloudInit: QemuCloudInitDisk{}.mapToStruct(*tmpCdRom)}
		}
	}

	tmpDisk := qemuDisk{}.mapToStruct(settings)
	if tmpDisk == nil {
		return nil
	}
	if tmpDisk.File == "" {
		return &QemuSataStorage{Disk: &QemuSataDisk{
			AsyncIO:    tmpDisk.AsyncIO,
			Backup:     tmpDisk.Backup,
			Bandwidth:  tmpDisk.Bandwidth,
			Cache:      tmpDisk.Cache,
			Discard:    tmpDisk.Discard,
			EmulateSSD: tmpDisk.EmulateSSD,
			Format:     tmpDisk.Format,
			Id:         tmpDisk.Id,
			Replicate:  tmpDisk.Replicate,
			Serial:     tmpDisk.Serial,
			Size:       tmpDisk.Size,
			Storage:    tmpDisk.Storage,
		}}
	}
	return &QemuSataStorage{Passthrough: &QemuSataPassthrough{
		AsyncIO:    tmpDisk.AsyncIO,
		Backup:     tmpDisk.Backup,
		Bandwidth:  tmpDisk.Bandwidth,
		Cache:      tmpDisk.Cache,
		Discard:    tmpDisk.Discard,
		EmulateSSD: tmpDisk.EmulateSSD,
		File:       tmpDisk.File,
		Replicate:  tmpDisk.Replicate,
		Serial:     tmpDisk.Serial,
		Size:       tmpDisk.Size,
	}}
}
