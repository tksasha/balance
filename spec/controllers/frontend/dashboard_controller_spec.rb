# frozen_string_literal: true

RSpec.describe Frontend::DashboardController, type: :controller do
  describe '#initialize_resource' do
    before { subject.send(:initialize_resource) }

    it { expect(subject.send(:resource)).to be_a(Frontend::Dashboard) }
  end
end
