# frozen_string_literal: true

module Frontend
  class Dashboard
    class CashesController < ApplicationController
      def update
        render :edit unless resource.update(resource_params)
      end

      private

      def resource
        @resource ||= Cash.find(params[:id])
      end

      def resource_params
        params.require(:cash).permit(:name, :formula)
      end

      def dashboard
        ::Frontend::Dashboard.new(params)
      end

      helper_method :dashboard
    end
  end
end
