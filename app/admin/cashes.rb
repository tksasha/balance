# frozen_string_literal: true

ActiveAdmin.register Cash do
  menu priority: 4, label: proc { I18n.t('active_admin.cashes') }

  permit_params :name
end
