# frozen_string_literal: true

ActiveAdmin.register Cash do
  menu label: proc { I18n.t('active_admin.cashes') }

  config.filters = false
  config.batch_actions = false

  actions :all, except: %i[destroy new create]
end
