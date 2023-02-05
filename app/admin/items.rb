# frozen_string_literal: true

ActiveAdmin.register Item do
  menu priority: 2, label: proc { I18n.t('active_admin.items') }

  actions :all, except: %i[destroy new create]

  permit_params %i[date]

  includes :category
end
