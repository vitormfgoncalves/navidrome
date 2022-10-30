import React from 'react'
import { Datagrid, TextField } from 'react-admin'
import { useMediaQuery } from '@material-ui/core'
import { SimpleList, List, NavButtons } from '../common'
import config from '../config'

const navStyle = {
  marginTop: '-15px',
  marginLeft: '15px',
  marginRight: '1em',
}

const TranscodingList = (props) => {
  const isXsmall = useMediaQuery((theme) => theme.breakpoints.down('xs'))
  return (
    <div>
      <div style={navStyle}>
        <NavButtons />
      </div>
      <List exporter={false} {...props}>
        {isXsmall ? (
          <SimpleList
            primaryText={(r) => r.name}
            secondaryText={(r) => `format: ${r.targetFormat}`}
            tertiaryText={(r) => r.defaultBitRate}
          />
        ) : (
          <Datagrid rowClick={config.enableTranscodingConfig ? 'edit' : 'show'}>
            <TextField source="name" />
            <TextField source="targetFormat" />
            <TextField source="defaultBitRate" />
            <TextField source="command" />
          </Datagrid>
        )}
      </List>
    </div>
  )
}

export default TranscodingList
