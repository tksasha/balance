# frozen_string_literal: true

RSpec.describe BaseController do
  describe '#dashboard' do
    before do
      allow(subject).to receive(:params).and_return(:params)

      allow(Frontend::Dashboard).to receive(:new).with(:params).and_return(:dashboard)
    end

    its(:dashboard) { is_expected.to eq :dashboard }

    its(:_helper_methods) { is_expected.to include :dashboard }
  end
end
