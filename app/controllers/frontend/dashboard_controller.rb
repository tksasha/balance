# frozen_string_literal: true

module Frontend
  class DashboardController < ApplicationController
    before_action :initialize_resource, only: :show

    private

    attr_reader :resource

    def initialize_resource
      @resource = ::Frontend::Dashboard.new
    end
  end
end
