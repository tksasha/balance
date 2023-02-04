# frozen_string_literal: true

ActiveAdmin.register Category do
  menu label: proc { I18n.t('active_admin.categories') }

  config.filters = false
  config.batch_actions = false

  actions :all, except: %i[destroy new create]
end
