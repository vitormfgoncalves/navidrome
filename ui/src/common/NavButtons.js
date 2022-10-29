import React from 'react'
import IconButton from '@material-ui/core/IconButton'
import ArrowBackIos from '@material-ui/icons/ChevronLeft'
import ArrowForwardIos from '@material-ui/icons/ChevronRight'
import { useHistory } from 'react-router-dom'
import { makeStyles } from '@material-ui/core/styles'

const useStyles = makeStyles(
  (theme) => ({
    button: {
      backgroundColor: 'rgba(255, 255, 255, 0.08);',
      border: 'none',
      borderRadius: 30,
      padding: '7px',
      '&:hover': {
        background: theme.palette.primary.dark,
      },
    },
    grpDiv: {
      marginRight: '-0.9em',
      marginTop: '30px',
      marginBottom: '9px',
      display: 'flex',
    },
    btDivider: {
      marginRight: '5px',
      display: 'flex',
    },
  }),
  { name: 'NavigationButtons' }
)

export const NavButtons = () => {
  const navigate = useHistory()
  const classes = useStyles()

  return (
    <div className={classes.grpDiv} id="navButtons">
      <IconButton className={classes.button} onClick={navigate.goBack}>
        <ArrowBackIos />
      </IconButton>
      <span className={classes.btDivider}></span>
      <IconButton className={classes.button} onClick={navigate.goForward}>
        <ArrowForwardIos />
      </IconButton>
    </div>
  )
}
