# frozen_string_literal: true

# rubocop:disable RSpec/HooksBeforeExamples
RSpec.shared_examples 'index' do |uri|
  describe "GET #{ uri }" do
    include_examples 'format', uri

    include_examples 'xhr?'

    before { get uri, xhr: xhr? }

    it { should render_template :index }

    it_behaves_like 'http_status', :ok

    it_behaves_like 'media_type'
  end
end

RSpec.shared_examples 'new' do |uri|
  describe "GET #{ uri }" do
    include_examples 'format', uri

    include_examples 'xhr?'

    before { get uri, xhr: xhr? }

    it { should render_template :new }

    it_behaves_like 'http_status', :ok

    it_behaves_like 'media_type'
  end
end

RSpec.shared_examples 'edit' do |uri|
  describe "GET #{ uri }" do
    include_examples 'format', uri

    include_examples 'xhr?'

    before { get uri, xhr: xhr? }

    it { should render_template :edit }

    it_behaves_like 'http_status', :ok

    it_behaves_like 'media_type'
  end
end

RSpec.shared_examples 'destroy' do |uri|
  describe "DELETE #{ uri }" do
    include_examples 'format', uri

    include_examples 'xhr?'

    before { delete uri, xhr: xhr? }

    it { success.call }

    it_behaves_like 'http_status', :ok

    it_behaves_like 'media_type'
  end
end

RSpec.shared_examples 'create' do |uri|
  describe "POST #{ uri }" do
    include_examples 'format', uri

    include_examples 'xhr?'

    context 'create with valid params' do
      before { post uri, params: valid_params, xhr: xhr? }

      it { success.call }

      it_behaves_like 'http_status', :created

      it_behaves_like 'media_type'
    end

    context 'create with invalid params' do
      before { post uri, params: invalid_params, xhr: xhr? }

      it { failure.call }

      it_behaves_like 'http_status', :unprocessable_entity

      it_behaves_like 'media_type'
    end
  end
end

RSpec.shared_examples 'update' do |uri|
  describe "PATCH #{ uri }" do
    include_examples 'format', uri

    include_examples 'xhr?'

    context 'update with valid params' do
      before { patch uri, params: valid_params, xhr: xhr? }

      it { success.call }

      it_behaves_like 'http_status', :ok

      it_behaves_like 'media_type'
    end

    context 'update with invalid params' do
      before { patch uri, params: invalid_params, xhr: xhr? }

      it { failure.call }

      it_behaves_like 'http_status', :unprocessable_entity

      it_behaves_like 'media_type'
    end
  end
end

RSpec.shared_context 'format' do |uri|
  let :format do
    #
    # '/items.json' -> 'json'
    #
    # '/items.js' -> 'js'
    #
    # '/items' -> 'html'
    #
    (md = /\.([[:alpha:]]+)$/.match(uri)) ? md[1] : 'html'
  end
end

RSpec.shared_context 'xhr?' do
  let :xhr? do
    format == 'js'
  end
end

RSpec.shared_examples 'http_status' do |status|
  context do
    subject { response }

    it { should have_http_status(status) }
  end
end

RSpec.shared_examples 'media_type' do
  context do
    subject { response }

    its :media_type do
      content_type = \
        case format
        when 'html'
          'text/html'
        when 'json'
          'application/json'
        when 'js'
          'text/javascript'
        end

      should eq content_type
    end
  end
end
# rubocop:enable RSpec/HooksBeforeExamples
