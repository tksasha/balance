require 'rails_helper'

class DummiesController
  include AbstractController::Callbacks
  include AbstractController::Helpers
  include ActsAsRESTController
end

class Dummy; end

RSpec.describe ActsAsRESTController do
  subject { DummiesController.new }

  its(:resource_class) { should eq Dummy }

  describe '#initialize_resource' do
    before { expect(Dummy).to receive(:new).and_return(:resource) }

    its(:initialize_resource) { should eq :resource }
  end

  describe '#build_resource' do
    before { expect(subject).to receive(:resource_params).and_return(:params) }

    before { expect(Dummy).to receive(:new).with(:params).and_return(:resource) }

    its(:build_resource) { should eq :resource }
  end

  describe '#resource' do
    before { expect(subject).to receive(:params).and_return({ id: 1 }) }

    before { expect(Dummy).to receive(:find).with(1).and_return(:resource) }

    its(:resource) { should eq :resource }
  end

  describe '#create' do
    let(:resource) { double }

    before { expect(subject).to receive(:resource).and_return(resource) }

    before { expect(resource).to receive(:save).and_return(false) }

    before { expect(subject).to receive(:render).with(:new) }

    it { expect { subject.create }.to_not raise_error }
  end

  describe '#update' do
    let(:resource) { double }

    before { expect(subject).to receive(:resource).and_return(resource) }

    before { expect(subject).to receive(:resource_params).and_return(:params) }

    before { expect(resource).to receive(:update).with(:params).and_return(false) }

    before { expect(subject).to receive(:render).with(:edit) }

    it { expect { subject.update }.to_not raise_error }
  end

  describe '#destroy' do
    let(:resource) { double }

    before { expect(subject).to receive(:resource).and_return(resource) }

    before { expect(resource).to receive(:destroy) }

    it { expect { subject.destroy }.to_not raise_error }
  end
end
