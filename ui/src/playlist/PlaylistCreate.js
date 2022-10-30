import React from 'react'
import {
  Create,
  SimpleForm,
  TextInput,
  BooleanInput,
  required,
  useTranslate,
  useRefresh,
  useNotify,
  useRedirect,
} from 'react-admin'
import { Title, NavButtons } from '../common'

const PlaylistCreate = (props) => {
  const { basePath } = props
  const refresh = useRefresh()
  const notify = useNotify()
  const redirect = useRedirect()
  const translate = useTranslate()
  const resourceName = translate('resources.playlist.name', { smart_count: 1 })
  const title = translate('ra.page.create', {
    name: `${resourceName}`,
  })

  const onSuccess = () => {
    notify('ra.notification.created', 'info', { smart_count: 1 })
    redirect('list', basePath)
    refresh()
  }

  const navStyle = {
    marginTop: '-15px',
    marginLeft: '15px',
    marginRight: '1em',
  }

  return (
    <>
      <div style={navStyle}>
        <NavButtons />
      </div>
      <Create
        title={<Title subTitle={title} />}
        {...props}
        onSuccess={onSuccess}
      >
        <SimpleForm redirect="list" variant={'outlined'}>
          <TextInput source="name" validate={required()} />
          <TextInput multiline source="comment" />
          <BooleanInput source="public" initialValue={true} />
        </SimpleForm>
      </Create>
    </>
  )
}

export default PlaylistCreate
