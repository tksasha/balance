# frozen_string_literal: true

ActiveAdmin.register Item do
  menu priority: 2, label: proc { I18n.t('active_admin.items') }

  config.filters = false
  config.batch_actions = false

  actions :all, except: %i[destroy new create]

  controller do
    def scoped_collection
      super.includes(:category)
    end
  end
end
