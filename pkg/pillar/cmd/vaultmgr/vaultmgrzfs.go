// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package vaultmgr

import (
	log "github.com/sirupsen/logrus"
)

const (
	zfsPath              = "/usr/sbin/chroot"
	defaultZpool         = "persist"
	defaultSecretDataset = defaultZpool + "/vault"
	zfsHostfsKeyFile     = "/containers/services/pillar/rootfs/var/run/TmpVaultDir2/protector.key"
	zfsKeyFile           = zfsKeyDir + "/protector.key"
	zfsKeyDir            = "/var/run/TmpVaultDir2"
)

func getCreateParams(vaultPath string) []string {
	args := []string{"/hostfs", "zfs", "create", "-o", "encryption=aes-256-gcm", "-o", "keylocation=file://" + zfsHostfsKeyFile, "-o", "keyformat=raw", vaultPath}
	return args
}

func getLoadKeyParams(vaultPath string) []string {
	args := []string{"/hostfs", "zfs", "load-key", vaultPath}
	return args
}

func getMountParams(vaultPath string) []string {
	args := []string{"/hostfs", "zfs", "mount", vaultPath}
	return args
}

func getKeyStatusParams(vaultPath string) []string {
	args := []string{"/hostfs", "zfs", "get", "keystatus", vaultPath}
	return args
}

//e.g. zfs load-key persist/vault followed by
//zfs mount persist/vault
func unlockZfsVault(vaultPath string) error {
	//prepare key in the staging file
	if err := stageKey(false, zfsKeyDir, zfsKeyFile); err != nil {
		return err
	}
	defer unstageKey(zfsKeyDir, zfsKeyFile)

	//zfs load-key
	args := getLoadKeyParams(vaultPath)
	if stdOut, stdErr, err := execCmd(zfsPath, args...); err != nil {
		log.Errorf("Error loading key for vault: %v, %s, %s",
			err, stdOut, stdErr)
		return err
	}
	//zfs mount
	args = getMountParams(vaultPath)
	if stdOut, stdErr, err := execCmd(zfsPath, args...); err != nil {
		log.Errorf("Error unlocking vault: %v, %s, %s", err, stdOut, stdErr)
		return err
	}
	return nil
}

//e.g. zfs create -o encryption=aes-256-gcm -o keylocation=file://tmp/raw.key -o keyformat=raw perist/vault
func createZfsVault(vaultPath string) error {
	//prepare key in the staging file
	if err := stageKey(false, zfsKeyDir, zfsKeyFile); err != nil {
		return err
	}
	defer unstageKey(zfsKeyDir, zfsKeyFile)
	args := getCreateParams(vaultPath)
	if stdOut, stdErr, err := execCmd(zfsPath, args...); err != nil {
		log.Errorf("Error creating zfs vault %s, error=%v, %s, %s",
			vaultPath, err, stdOut, stdErr)
		return err
	}
	log.Infof("Created new vault %s", vaultPath)
	return nil
}

//e.g. zfs get keystatus persist/vault
func checkKeyStatus(vaultPath string) error {
	args := getKeyStatusParams(vaultPath)
	if stdOut, stdErr, err := execCmd(zfsPath, args...); err != nil {
		log.Debugf("keystatus query for %s results in error=%v, %s, %s",
			vaultPath, err, stdOut, stdErr)
		return err
	}
	return nil
}

func setupZfsVault(vaultPath string) error {
	//zfs get keystatus returns success as long as vaultPath is a dataset,
	//(even if not mounted yet), so use it to check dataset presence
	if err := checkKeyStatus(vaultPath); err == nil {
		//present, call unlock
		return unlockZfsVault(vaultPath)
	}
	//try creating the dataset
	return createZfsVault(vaultPath)
}